package I18n

var I18nDict_tg_JP map[uint16]string = map[uint16]string{
	ACME_FailedToGetCommand:             "ACME コマンドの解析に失敗した。",
	ACME_FailedToSeek:                   "ファイルのシークに失敗したんで、この ACME ファイルが無効かもしれん。",
	ACME_StructureErrorNotice:           "ファイルの内容が無効だ。",
	ACME_UnknownCommand:                 "このファイルには未知な ACME コマンドがある。",
	Auth_BackendError:                   "バックエンドエラー",
	Auth_FailedToRequestEntry:           "サーバーの入り方が貰えないんで、サーバーのレベルレギュレーションを締めて、パスワードをチェックしてね。",
	Auth_HelperNotCreated:               "ヘルパーユーザーがまだないんで、FastBuilder ユーザーセンターで作ってね。",
	Auth_InvalidFBVersion:               "このバーションはアウトデートなので更新して。",
	Auth_InvalidHelperUsername:          "ヘルパーユーザーがまだ名付かれていないんで、FastBuilder ユーザーセンターで名付けてください。",
	Auth_InvalidToken:                   "ログイントークンが無効。",
	Auth_InvalidUser:                    "ユーザー無効。",
	Auth_ServerNotFound:                 "指定されたサーバーが見つかれん。開放状態を確認したあとまたやってみて。",
	Auth_UnauthorizedRentalServerNumber: "未認証のサーバー番号だった。FastBuilder ユーザーセンターでアドしてください。",
	Auth_UserCombined:                   "指定されたユーザーはもうほかのユーザーに合体されたんで新アカウントでログインしてください。",
	Auth_FailedToRequestEntry_TryAgain:  "サーバーインフォの取得に失敗した。あとでもう一度試してね。",
	BDump_Author:                        "作者",
	BDump_EarlyEOFRightWhenOpening:      "ファイルエンドが来るのは早すぎたんで読めんぞい。それがもうこわれたかも。",
	BDump_FailedToGetCmd1:               "cmd[pos:0] のアーギュメントが読めん。ファイルがもう壊れたかも。",
	BDump_FailedToGetCmd2:               "cmd[pos1] のアーギュメントが読めん。ファイルがもう壊れたかも。",
	BDump_FailedToGetCmd4:               "cmd[pos2] のアーギュメントが読めん。ファイルがもう壊れたかも。",
	BDump_FailedToGetCmd6:               "cmd[pos3] のアーギュメントが読めん。ファイルがもう壊れたかも。",
	BDump_FailedToGetCmd7_0:             "cmd[pos4] のアーギュメントが読めん。ファイルがもう壊れたかも。",
	BDump_FailedToGetCmd7_1:             "cmd[pos5] のアーギュメントが読めん。ファイルがもう壊れたかも。",
	BDump_FailedToGetCmd10:              "cmd[pos6] のアーギュメントが読めん。ファイルがもう壊れたかも。",
	BDump_FailedToGetCmd11:              "cmd[pos7] のアーギュメントが読めん。ファイルがもう壊れたかも。",
	BDump_FailedToGetCmd12:              "cmd[pos8] のアーギュメントが読めん。ファイルがもう壊れたかも。",
	BDump_FailedToGetConstructCmd:       "BDump コマンドが読めん。ファイルがもう壊れたかも。",
	BDump_FailedToReadAuthorInfo:        "作者情報が読めない。ファイルがもう壊れたかも。",
	BDump_FileNotSigned:                 "このファイルはサインされてない。",
	BDump_FileSigned:                    "このファイルはサインされた。サイン元: %s",
	BDump_NotBDX_Invheader:              "このファイルは bdx ファイルじゃないんだ。(無効なヘッダー)",
	BDump_NotBDX_Invinnerheader:         "ファイルが bdx ファイルじゃない。 (無効な内部ヘッダー)",
	BDump_SignedVerifying:               "このファイルはサインされた、それをいまチェックしてる...",
	BDump_VerificationFailedFor:         "%v のためにファイルのサインをチェック出来ん。",
	BDump_Warn_Reserved:                 "警告: BDump/Import: 保留されたコマンドがいま使用された\n",
	CommandNotFound:                     "未知なコマンド。",
	ConnectionEstablished:               "サーバーへ接続した。",
	Copyright_Notice_Bouldev:            "Copyright (c) FastBuilder DevGroup, Bouldev 2024",
	Copyright_Notice_Contrib:            "コントリビュータ: Ruphane, CAIMEO, CMA2401PT",
	Crashed_No_Connection:               "長い時間をかけても接続できなかった。",
	Crashed_OS_Windows:                  "ENTER を押して終了。",
	Crashed_StackDump_And_Error:         "スタックダンプ (Stack dump) が上に見える。エラーは: ",
	Crashed_Tip:                         "ヤバいぞい！FastBuilder Phoenix がクラッシューしたよ！",
	CurrentDefaultDelayMode:             "今のデフォルトディレーモード",
	CurrentTasks:                        "進行中のタスク:",
	DelayModeSet:                        "ディレーモードが設定された",
	DelayModeSet_DelayAuto:              "ディレーを自動に %d に設定した。",
	DelayModeSet_ThresholdAuto:          "ディレーのしきい値を自動に  %d に設定した．",
	DelaySet:                            "ディレーが設定された．",
	DelaySetUnavailableUnderNoneMode:    "[delay set] コマンドがディレーモードが none の時に使えない。",
	DelayThreshold_OnlyDiscrete:         "ディレーのしきい値がディレーモードが discrete の時だけに使える。",
	DelayThreshold_Set:                  "ディレーのしきい値を %d にした。",
	ERRORStr:                            "エラー",
	EnterPasswordForFBUC:                "FastBuilder ユーザーセンターのパスワードは: ",
	Enter_FBUC_Username:                 "FastBuilder ユーザーセンターのログイン名は: ",
	Enter_Rental_Server_Code:            "サーバー番号は: ",
	Enter_Rental_Server_Password:        "サーバーのパスワードはなに？ (パスワードがなければ直接 ENTER をおす。入力した内容がみえん): ",
	ErrorIgnored:                        "エラーを無視した。",
	Error_MapY_Exceed:                   "3DMap で, MapY が必ず [20~255] の範囲にしてください (入力した方 = %v)",
	FBUC_LoginFailed:                    "FastBuilder ユーザーセンターのログイン名やパスワードが間違えた。",
	FBUC_Token_ErrOnCreate:              "トークンファイルが作成できん: ",
	FBUC_Token_ErrOnGen:                 "臨時トークンを作成出来なかった",
	FBUC_Token_ErrOnRemove:              "トークンファイルを削除できなかった: %v",
	FBUC_Token_ErrOnSave:                "トークンを保存できなかった: ",
	FileCorruptedError:                  "ファイルが壊れた",
	Get_Warning:                         "",
	IgnoredStr:                          "無視した",
	InvalidFileError:                    "無効なファイル。",
	InvalidPosition:                     "働くポジションを取得できなかった。 (無視でOK)",
	Lang_Config_ErrOnCreate:             "言語設定ファイルを作成できなかった: %v",
	Lang_Config_ErrOnSave:               "言語設定を保存できなかった: %v",
	LanguageName:                        "日本語（タメ語）",
	LanguageUpdated:                     "言語設定を変更した。",
	Logout_Done:                         "FastBuilder ユーザーセンターからログアウトした。",
	Menu_BackButton:                     "< Back", // REMOVED FEATURES NO TRANSLATION
	Menu_Cancel:                         "Cancel", // FOR THEM
	Menu_CurrentPath:                    "Current path",
	Menu_ExcludeCommandsOption:          "Exclude Commands",
	Menu_GetEndPos:                      "getEndPos",
	Menu_GetPos:                         "getPos",
	Menu_InvalidateCommandsOption:       "Invalidate Commands",
	Menu_Quit:                           "Quit Program",
	Menu_StrictModeOption:               "Strict Mode",
	NotAnACMEFile:                       "このファイルは ACME ストラクチャーファイルじゃなかった。",
	Notice_CheckUpdate:                  "更新を検査中、待っていてね...",
	Notice_iSH_Location_Service:         "iSH に居るんで、生きることを保持にはロケーションサービスが必要。位置情報が利用されんで、いつでもそれを締めてもいい。",
	Notice_OK:                           "完成\n",
	Notice_UpdateAvailable:              "新しいバーション (%s) の PhoenixBuilder がリリースされた。\n",
	Notice_UpdateNotice:                 "更新して。\n",
	Notice_ZLIB_CVE:                     "今使われてる zlib のバーション (%s) は古すぎていくつかの CVE バルネラビリティーがもう確認された、それを更新した方がいい。",
	Notify_NeedOp:                       "OP レベルが PhoenixBuilder の正常なかどうには必要。",
	Notify_TurnOnCmdFeedBack:            "PhoenixBuilder の正常な稼働には gamerule sendcommandfeedback を true にするのが必要なので、もう自動に　gamerule sendcommandfeedback を true にさせたんで、必要があればそのうち書き直してね。",
	Omega_WaitingForOP:                  "Omega System が OP レベルの取得を待っている...",
	Omega_Enabled:                       "Omega System が使われてる!",
	OpPrivilegeNotGrantedForOperation:   "OP レベルが PhoenixBuilder の正常なかどうには必要なので、OP にしたらまた試して。",
	Parsing_UnterminatedEscape:          "終わらず escape",
	Parsing_UnterminatedQuotedString:    "終わらず quoted string",
	PositionGot:                         "位置ゲット！",
	PositionGot_End:                     "終点位置ゲット！",
	PositionSet:                         "位置をセットした！",
	PositionSet_End:                     "終点位置をセットした",
	QuitCorrectly:                       "プログラムが正常に終了した。",
	Sch_FailedToResolve:                 "ファイルを読めなかった",
	SelectLanguageOnConsole:             "コンソールで新言語を選べて。",
	ServerCodeTrans:                     "サーバー",
	SimpleParser_Int_ParsingFailed:      "アーギュメント読み込み器: 整数アーギュメントを読めなかった。",
	SimpleParser_InvEnum:                "アーギュメント読み込み器: 予期せず選択肢、予期した選択肢は: %s.",
	SimpleParser_Invalid_decider:        "アーギュメント読み込み器: 無効な選択肢だった",
	SimpleParser_Too_few_args:           "アーギュメント読み込み器: アーギュメント少なさ過ぎる",
	Special_Startup:                     "日本語（タメ語）を使ってる\n",
	TaskCreated:                         "タスクを作った",
	TaskDisplayModeSet:                  "タスク状態表示モードを %s にした．",
	TaskFailedToParseCommand:            "コマンド %v を読めなかった。",
	TaskNotFoundMessage:                 "教えたタスク番号で進行してるタスクを見つかれんぞ。",
	TaskPausedNotice:                    "[タスク %d] - 一時停止",
	TaskResumedNotice:                   "[タスク %d] - 回復",
	TaskStateLine:                       "ID %d - コマンドライン:\"%s\", 状態: %s, ディレー: %d, ディレーモード: %s, ディレーしきい値: %d",
	TaskStoppedNotice:                   "[タスク %d] - 停止",
	TaskTTeIuKoto:                       "タスク",
	TaskTotalCount:                      "総計: %d",
	TaskTypeCalculating:                 "計算中",
	TaskTypeDied:                        "停止",
	TaskTypePaused:                      "一時停止",
	TaskTypeRunning:                     "進行中",
	TaskTypeSpecialTaskBreaking:         "スペシャルタスク:停止中",
	TaskTypeSwitchedTo:                  "タスク作成タイプを %s にした",
	TaskTypeUnknown:                     "未知",
	Task_D_NothingGenerated:             "[タスク %d] なにも建造されなかった",
	Task_DelaySet:                       "[タスク %d] - ディレーを %d にした",
	Task_ResumeBuildFrom:                "ブロック数 %d から建造を再開する",
	Task_SetDelay_Unavailable:           "[setdelay] がディレーモードが none の時に使えん。",
	Task_Summary_1:                      "[タスク %d] %v 影響されたブロック",
	Task_Summary_2:                      "[タスク %d] %v 秒をかけた",
	Task_Summary_3:                      "[タスク %d] 平均速度: %v ブロック/秒",
	UnsupportedACMEVersion:              "このバーションの ACME ストラクチャーファイルは支援されていない。バーション 1.2 だけが支援されてる。",
	Warning_ACME_Deprecated:             "警告 - `acme' は非推奨となり、削除されましたので、代わりにBDX形式に移行していただくようお願いいたします。詳細については、https://github.com/LNSSPsd/PhoenixBuilder/issues/313 をご参照くださいませ。",
	Warning_Schem_Deprecated:            "警告 - `schem' は非推奨となり、削除されましたので、代わりにBDX形式に移行していただくようお願いいたします。詳細については、https://github.com/LNSSPsd/PhoenixBuilder/issues/313 をご参照くださいませ。",
	Warning_UserHomeDir:                 "警告 - ユーザーのホームフォルダーが見つかれん。homedir=\".\";にした。\n",
}
