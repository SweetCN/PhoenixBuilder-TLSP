package uqHolder

import (
	"bufio"
	"bytes"
	"encoding/gob"
	"encoding/json"
	"fmt"
	"phoenixbuilder/minecraft/protocol"
	"phoenixbuilder/minecraft/protocol/packet"
	"time"

	"github.com/go-gl/mathgl/mgl32"
	"github.com/google/uuid"
)

var Version = [3]byte{0, 0, 1}

type Player struct {
	UUID                    uuid.UUID
	EntityUniqueID          int64
	Username                string
	PlatformChatID          string
	BuildPlatform           int32
	SkinID                  string
	PropertiesFlag          uint32
	CommandPermissionLevel  uint32
	ActionPermissions       uint32
	OPPermissionLevel       uint32
	CustomStoredPermissions uint32
	// only when the player can be seen by bot
	EntityRuntimeID uint64
	Entity          *Entity
	// PlayerUniqueID is a unique identifier of the player. It appears it is not required to fill this field
	// out with a correct value. Simply writing 0 seems to work.
	PlayerUniqueID int64
}

type PosRepresent struct {
	Position       mgl32.Vec3
	Velocity       mgl32.Vec3
	Pitch          float32
	Yaw            float32
	HeadYaw        float32
	LastUpdateTick uint64
	Rotation       mgl32.Vec3
	MaskedRotation mgl32.Vec3
}

type Entity struct {
	RuntimeID        uint64
	Attributes       []protocol.Attribute
	Metadata         map[uint32]interface{}
	Slots            map[byte]*Equipment
	LastPacketSlot   byte
	OutOfRangeAtTick uint64
	IsPlayer         bool

	LastUpdateTick uint64
	LastPosInfo    PosRepresent

	UniqueID    int64
	EntityType  string
	EntityLinks []protocol.EntityLink
}

type Equipment struct {
	NewItem  protocol.ItemInstance
	Slot     byte
	WindowID byte
}

type GameRule struct {
	CanBeModifiedByPlayer bool
	Value                 interface{}
}
type UQHolder struct {
	VERSION             string
	BotRuntimeID        uint64
	CompressThreshold   uint16
	CurrentTick         uint64
	InventorySlot       map[uint32]protocol.ItemInstance
	playersByUUID       map[[16]byte]*Player
	PlayersByEntityID   map[int64]*Player
	EntitiesByRuntimeID map[uint64]*Entity
	entitiesByUniqueID  map[int64]*Entity
	Time                int32
	DayTime             int32
	DayTimePercent      float32
	WorldSpawnPosition  map[int32]protocol.BlockPos
	BotSpawnPosition    map[int32]protocol.BlockPos
	Difficulty          uint32
	CommandsEnabled     bool
	GameRules           map[string]*GameRule
	InventoryContent    map[uint32][]protocol.ItemInstance
	PlayerHotBar        packet.PlayerHotBar
	AvailableCommands   packet.AvailableCommands
	BotOnGround         bool
	BotHealth           int32
	CommandRelatedEnums []*packet.UpdateSoftEnum
}

func NewUQHolder(BotRuntimeID uint64) *UQHolder {
	uq := &UQHolder{
		VERSION:             fmt.Sprintf("%d.%d.%d", Version[0], Version[1], Version[2]),
		BotRuntimeID:        BotRuntimeID,
		InventorySlot:       map[uint32]protocol.ItemInstance{},
		playersByUUID:       map[[16]byte]*Player{},
		PlayersByEntityID:   map[int64]*Player{},
		WorldSpawnPosition:  map[int32]protocol.BlockPos{},
		BotSpawnPosition:    map[int32]protocol.BlockPos{},
		EntitiesByRuntimeID: map[uint64]*Entity{},
		entitiesByUniqueID:  map[int64]*Entity{},
		GameRules:           map[string]*GameRule{},
		InventoryContent:    map[uint32][]protocol.ItemInstance{},
		CommandRelatedEnums: make([]*packet.UpdateSoftEnum, 0),
	}
	go func() {
		t := time.NewTicker(50 * time.Millisecond)
		for {
			<-t.C
			uq.CurrentTick++
		}
	}()
	return uq
}

func (uq *UQHolder) UpdateTick(tick uint64) {
	uq.CurrentTick = tick
}

func (uq *UQHolder) Marshal() []byte {
	buf := bytes.NewBuffer([]byte{Version[0], Version[1], Version[2]})
	err := gob.NewEncoder(buf).Encode(uq)
	if err != nil {
		panic(err)
	}
	return buf.Bytes()
}

func IsCapable(bs []byte) error {
	if len(bs) < 3 {
		return fmt.Errorf("version length error")
	}
	if bs[0] != Version[0] {
		return fmt.Errorf("version MAJOR mismatch (local=%v,remote=%v)", Version[0], bs[0])
	}
	if bs[1] != Version[1] {
		return fmt.Errorf("version MINOR mismatch (local=%v,remote=%v)", Version[1], bs[1])
	}
	if bs[2] != Version[2] {
		return fmt.Errorf("version Patch mismatch (local=%v,remote=%v)", Version[2], bs[2])
	}
	return nil
}

func (uq *UQHolder) UnMarshal(bs []byte) error {
	if err := IsCapable(bs); err != nil {
		return err
	}
	buf := bytes.NewBuffer(bs[3:])
	bufio.NewReader(buf)
	err := gob.NewDecoder(buf).Decode(uq)
	if err != nil {
		return err
	}
	for _, entity := range uq.EntitiesByRuntimeID {
		uq.entitiesByUniqueID[entity.UniqueID] = entity
	}
	for _, player := range uq.PlayersByEntityID {
		uq.playersByUUID[player.UUID] = player
		if player.EntityRuntimeID != 0 {
			if e, ok := uq.EntitiesByRuntimeID[player.EntityRuntimeID]; ok {
				player.Entity = e
			}
		}
	}
	return nil
}

func (uq *UQHolder) GetEntityByRuntimeID(EntityRuntimeID uint64) *Entity {
	var e *Entity
	if _e, ok := uq.EntitiesByRuntimeID[EntityRuntimeID]; !ok {
		e = &Entity{
			RuntimeID:      EntityRuntimeID,
			LastPacketSlot: 255,
			Slots:          map[byte]*Equipment{},
			LastUpdateTick: uq.CurrentTick,
			UniqueID:       0,
		}
		uq.EntitiesByRuntimeID[EntityRuntimeID] = e
	} else {
		e = _e
	}
	return e
}

func (uq *UQHolder) Update(pk packet.Packet) {
	defer func() {
		r := recover()
		if r != nil {
			fmt.Println("UQHolder Update Error: ", r)
		}
	}()
	switch p := pk.(type) {
	case *packet.NetworkSettings:
		uq.CompressThreshold = p.CompressionThreshold
	case *packet.InventorySlot:
		uq.InventorySlot[p.Slot] = p.NewItem
	case *packet.PlayerList:
		if p.ActionType == packet.PlayerListActionAdd {
			for _, e := range p.Entries {
				player := &Player{
					UUID:           e.UUID,
					EntityUniqueID: e.EntityUniqueID,
					Username:       e.Username,
					PlatformChatID: e.PlatformChatID,
					BuildPlatform:  e.BuildPlatform,
					SkinID:         e.Skin.SkinID,
				}
				uq.playersByUUID[e.UUID] = player
				uq.PlayersByEntityID[e.EntityUniqueID] = player
			}
		} else {
			for _, e := range p.Entries {
				delete(uq.playersByUUID, e.UUID)
			}
		}
	case *packet.AdventureSettings:
		player := uq.PlayersByEntityID[p.PlayerUniqueID]
		player.PropertiesFlag = p.Flags
		player.CommandPermissionLevel = p.CommandPermissionLevel
		player.ActionPermissions = p.ActionPermissions
		player.OPPermissionLevel = p.PermissionLevel
		player.CustomStoredPermissions = p.CustomStoredPermissions
	case *packet.SetTime:
		uq.Time = p.Time
		uq.DayTime = p.Time % 24000
		uq.DayTimePercent = float32(uq.DayTime) / 24000.0
	case *packet.SetDifficulty:
		uq.Difficulty = p.Difficulty
	case *packet.SetCommandsEnabled:
		uq.CommandsEnabled = p.Enabled
	case *packet.UpdateAttributes:
		e := uq.GetEntityByRuntimeID(p.EntityRuntimeID)
		e.LastUpdateTick = p.Tick
		e.Attributes = p.Attributes
		uq.UpdateTick(p.Tick)
	case *packet.GameRulesChanged:
		for _, r := range p.GameRules {
			uq.GameRules[r.Name] = &GameRule{
				CanBeModifiedByPlayer: r.CanBeModifiedByPlayer,
				Value:                 r.Value,
			}
		}
	case *packet.InventoryContent:
		uq.InventoryContent[p.WindowID] = p.Content

	case *packet.AvailableCommands:
		uq.AvailableCommands = *p

	case *packet.SetActorData:
		e := uq.GetEntityByRuntimeID(p.EntityRuntimeID)
		e.LastUpdateTick = p.Tick
		e.Metadata = p.EntityMetadata
		uq.UpdateTick(p.Tick)

	case *packet.MovePlayer:
		if p.EntityRuntimeID == uq.BotRuntimeID {
			uq.BotOnGround = p.OnGround
		}
		e := uq.GetEntityByRuntimeID(p.EntityRuntimeID)
		e.LastPosInfo = PosRepresent{
			Position:       p.Position,
			Pitch:          p.Pitch,
			Yaw:            p.Yaw,
			HeadYaw:        p.HeadYaw,
			LastUpdateTick: p.Tick,
		}
		e.LastUpdateTick = p.Tick
		uq.UpdateTick(p.Tick)
	case *packet.CorrectPlayerMovePrediction:
		uq.GetEntityByRuntimeID(uq.BotRuntimeID).LastPosInfo.Position = p.Position
		uq.GetEntityByRuntimeID(uq.BotRuntimeID).LastPosInfo.LastUpdateTick = p.Tick
		uq.GetEntityByRuntimeID(uq.BotRuntimeID).LastUpdateTick = p.Tick
		uq.BotOnGround = p.OnGround
		uq.UpdateTick(p.Tick)

	case *packet.AddPlayer:
		player := uq.PlayersByEntityID[p.EntityUniqueID]
		entity := uq.GetEntityByRuntimeID(p.EntityRuntimeID)
		entity.IsPlayer = true
		entity.LastUpdateTick = uq.CurrentTick
		player.Entity = entity
		entity.LastUpdateTick = uq.CurrentTick
		entity.LastPosInfo.LastUpdateTick = uq.CurrentTick
		entity.LastPosInfo.Position = p.Position
		entity.LastPosInfo.Pitch = p.Pitch
		entity.LastPosInfo.Yaw = p.Yaw
		entity.LastPosInfo.HeadYaw = p.HeadYaw
		player.PropertiesFlag = p.Flags
		player.CommandPermissionLevel = p.CommandPermissionLevel
		player.ActionPermissions = p.ActionPermissions
		player.OPPermissionLevel = p.PermissionLevel
		player.CustomStoredPermissions = p.CustomStoredPermissions
		player.PlayerUniqueID = p.PlayerUniqueID

	case *packet.MobEquipment:
		entity := uq.GetEntityByRuntimeID(p.EntityRuntimeID)
		entity.Slots[p.InventorySlot] = &Equipment{
			Slot:     p.InventorySlot,
			NewItem:  p.NewItem,
			WindowID: p.WindowID,
		}
		entity.LastUpdateTick = uq.CurrentTick
		entity.LastPacketSlot = p.InventorySlot
	case *packet.SetHealth:
		uq.BotHealth = p.Health
	case *packet.UpdateSoftEnum:
		uq.CommandRelatedEnums = append(uq.CommandRelatedEnums, p)
	case *packet.AddActor:
		entity := uq.GetEntityByRuntimeID(p.EntityRuntimeID)
		entity.IsPlayer = false
		entity.UniqueID = p.EntityUniqueID
		uq.entitiesByUniqueID[p.EntityUniqueID] = entity
		entity.EntityType = p.EntityType
		entity.LastUpdateTick = uq.CurrentTick
		entity.LastPosInfo.LastUpdateTick = uq.CurrentTick
		entity.LastPosInfo.Position = p.Position
		entity.LastPosInfo.Velocity = p.Velocity
		entity.LastPosInfo.Pitch = p.Pitch
		entity.LastPosInfo.Yaw = p.Yaw
		entity.LastPosInfo.HeadYaw = p.Yaw
		entity.Attributes = p.Attributes
		entity.Metadata = p.EntityMetadata
		entity.EntityLinks = p.EntityLinks

	case *packet.RemoveActor:
		if entity, ok := uq.entitiesByUniqueID[p.EntityUniqueID]; ok {
			rtID := entity.RuntimeID
			if !entity.IsPlayer {
				if _, ok := uq.EntitiesByRuntimeID[rtID]; ok {
					delete(uq.EntitiesByRuntimeID, rtID)
				}
				delete(uq.entitiesByUniqueID, p.EntityUniqueID)
			}
		}
	case *packet.MoveActorDelta:
		entity := uq.GetEntityByRuntimeID(p.EntityRuntimeID)
		entity.LastPosInfo.LastUpdateTick = uq.CurrentTick
		entity.LastPosInfo.Position = p.Position
		entity.LastPosInfo.Rotation = p.Rotation
		if x := p.Rotation.X(); x != 0 {
			entity.LastPosInfo.MaskedRotation[0] = x
		}
		if y := p.Rotation.Y(); y != 0 {
			entity.LastPosInfo.MaskedRotation[1] = y
		}
		if z := p.Rotation.Z(); z != 0 {
			entity.LastPosInfo.MaskedRotation[2] = z
		}

	case *packet.SetActorMotion:
		entity := uq.GetEntityByRuntimeID(p.EntityRuntimeID)
		entity.LastPosInfo.LastUpdateTick = uq.CurrentTick
		entity.LastPosInfo.Velocity = p.Velocity

	// not fully supported
	case *packet.Respawn:
		if p.EntityRuntimeID == 0 {
			uq.GetEntityByRuntimeID(uq.BotRuntimeID).LastPosInfo.Position = p.Position
		} else {
			if marshal, err := json.Marshal(pk); err == nil {
				fmt.Println("Respawn Data ignored: ", string(marshal))
			}

		}
	// not fully supported, large amount of data
	case *packet.LevelEvent:

	// meaning not clear
	case *packet.SetSpawnPosition:
		if p.SpawnType == packet.SpawnTypePlayer {
			uq.BotSpawnPosition[p.Dimension] = p.Position
			uq.WorldSpawnPosition[p.Dimension] = p.SpawnPosition // not sure
		} else {
			uq.BotSpawnPosition[p.Dimension] = p.Position // not sure
			uq.WorldSpawnPosition[p.Dimension] = p.SpawnPosition
		}
	// meaning not clear
	case *packet.PlayerHotBar:
		uq.PlayerHotBar = *p
	// not supported, plan to support
	case *packet.InventoryTransaction:
	// not supported, plan to support
	case *packet.ActorEvent:
	// no plan to support the followings
	case *packet.LevelChunk:
	case *packet.NetworkChunkPublisherUpdate:
	case *packet.BiomeDefinitionList:
	case *packet.AvailableActorIdentifiers:
	case *packet.CraftingData:
	case *packet.ChunkRadiusUpdated:
	case *packet.LevelSoundEvent:
	case *packet.Animate:
	// no need to support
	case *packet.PlayStatus:
	// no need to support
	case *packet.PyRpc:
		//not handled
		//default:
		//	marshal, err := json.Marshal(pk)
		//	if err != nil {
		//		println(err)
		//	} else {
		//		jsonStr := string(marshal)
		//		if len(jsonStr) < 300 {
		//			fmt.Println(pk.ID(), " : ", jsonStr)
		//		} else {
		//			fmt.Println(pk.ID(), " : ", jsonStr[:300])
		//		}
		//	}
	}
}

//func main() {
//	TypePool := packet.NewPool()
//	fp, err := os.OpenFile("dump.gob", os.O_RDONLY, 0755)
//	if err != nil {
//		panic(err)
//	}
//	cachedBytes := make([][]byte, 0)
//	err = gob.NewDecoder(fp).Decode(&cachedBytes)
//	if err != nil {
//		panic(err)
//	}
//	holder := NewUQHolder(0)
//	paddingByte := []byte{}
//	safeDecode := func(pktByte []byte) (pkt packet.Packet) {
//		pktID := uint32(pktByte[0])
//		defer func() {
//			if r := recover(); r != nil {
//				fmt.Println(pktID, "decode fail ", pkt)
//			}
//			return
//		}()
//		pkt = TypePool[pktID]()
//		pkt.Unmarshal(protocol.NewReader(bytes.NewReader(
//			bytes.Join([][]byte{pktByte[1:], paddingByte}, []byte{}),
//		), 0))
//		return
//	}
//	for _, pktByte := range cachedBytes {
//      if len(pktByte)==0 {
//			continue
//		}
//		pkt := safeDecode(pktByte)
//		if pkt != nil {
//			holder.Update(pkt)
//		}
//	}
//}