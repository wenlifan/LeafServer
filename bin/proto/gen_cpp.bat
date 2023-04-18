@echo off

@REM echo convert client\Enum.proto...
@REM protoc.exe --proto_path=..\client --cpp_out=..\..\dsgame\Source\DSGame\Network\HandlerClient\Proto      Enum.proto
@REM
@REM echo convert client\Struct.proto...
@REM protoc.exe --proto_path=..\client --cpp_out=..\..\dsgame\Source\DSGame\Network\HandlerClient\Proto      Struct.proto
@REM
@REM
@REM echo convert client\PreLobby.proto...
@REM protoc.exe --proto_path=..\client --cpp_out=..\..\dsgame\Source\DSGame\Network\HandlerClient\Proto      PreLobby.proto
@REM
@REM echo convert client\Lobby.proto...
@REM protoc.exe --proto_path=..\client --cpp_out=..\..\dsgame\Source\DSGame\Network\HandlerClient\Proto      Lobby.proto
@REM
@REM echo convert client\Match.proto...
@REM protoc.exe --proto_path=..\client --cpp_out=..\..\dsgame\Source\DSGame\Network\HandlerClient\Proto      Match.proto
@REM
@REM echo convert client\ClientCommon.proto...
@REM protoc.exe --proto_path=..\client --cpp_out=..\..\dsgame\Source\DSGame\Network\HandlerClient\Proto      ClientCommon.proto
@REM
@REM echo convert client\Upgrade.proto...
@REM protoc.exe --proto_path=..\client --cpp_out=..\..\dsgame\Source\DSGame\Network\HandlerClient\Proto      Upgrade.proto
@REM
@REM echo convert client\Friend.proto...
@REM protoc.exe --proto_path=..\client --cpp_out=..\..\dsgame\Source\DSGame\Network\HandlerClient\Proto      Friend.proto
@REM
@REM echo convert server\Level.proto...
@REM protoc.exe --proto_path=..\server --cpp_out=..\..\dsgame\Source\DSGame\Network\HandlerServer\Proto      Level.proto
@REM
@REM echo convert client\Team.proto...
@REM protoc.exe --proto_path=..\client --cpp_out=..\..\dsgame\Source\DSGame\Network\HandlerClient\Proto      Team.proto
@REM
@REM echo convert client\Mail.proto...
@REM protoc.exe --proto_path=..\client --cpp_out=..\..\dsgame\Source\DSGame\Network\HandlerClient\Proto      Mail.proto
@REM
@REM echo convert client\Chat.proto...
@REM protoc.exe --proto_path=..\client --cpp_out=..\..\dsgame\Source\DSGame\Network\HandlerClient\Proto      Chat.proto

protoc.exe --proto_path=.\share --go_out=..\..\vendor Base.proto
protoc.exe --proto_path=.\client --go_out=..\..\vendor Struct.proto
protoc.exe --proto_path=.\client --go_out=..\..\vendor Enum.proto
protoc.exe --proto_path=.\client --go_out=..\..\vendor PreLobby.proto

echo finish!!!
pause