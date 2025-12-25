package main

import (
	_ "github.com/zzy-rabbit/patrol/data/dao"
	_ "github.com/zzy-rabbit/patrol/hardware/nfc"
	_ "github.com/zzy-rabbit/patrol/hardware/qrcode"
	_ "github.com/zzy-rabbit/patrol/logic/config"
	_ "github.com/zzy-rabbit/patrol/logic/executor"
	_ "github.com/zzy-rabbit/patrol/logic/trigger"
	_ "github.com/zzy-rabbit/patrol/protocol/http"
)
