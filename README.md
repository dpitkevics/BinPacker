# BinPacker

3D Bin packing solution written in Go Lang

Compile for linux on Mac OS:
`GOOS=linux GOARCH=386 go build -o BoxPacker-Linux github.com/dpitkevics/BoxPacker`

Compile for Mac OS on Mac OS:
`go build github.com/dpitkevics/BoxPacker`

Example on how to use this lib is found in this file:
`https://github.com/dpitkevics/BinPacker/blob/master/src/github.com/dpitkevics/BoxPacker/main.go`

Sample app call:  
`./BoxPacker '[{"Reference":"#72","OuterLength":24,"OuterWidth":18,"OuterHeight":24,"EmptyWeight":75,"InnerLength":24,"InnerWidth":18,"InnerHeight":24,"MaxWeight":99999}]' '[{"Identifier":"58eb1fed37c8b","Description":"My Test Item","Length":11.81,"Width":9.84,"Height":7.87,"Weight":3.52}]'`

Will return this response:  
`[{"Box":{"Reference":"#72","OuterLength":24,"OuterWidth":18,"OuterHeight":24,"EmptyWeight":75,"InnerLength":24,"InnerWidth":18,"InnerHeight":24,"InnerVolume":10368,"MaxWeight":99999},"Items":[{"Identifier":"58eb1fed37c8b","Description":"My Test Item","Length":11.81,"Width":9.84,"Height":7.87,"Weight":3.52,"Volume":914.5758480000001}],"Weight":0,"RemainingWidth":18,"RemainingLength":12.19,"RemainingHeight":24,"RemainingWeight":99920.48,"UsedWidth":9.84,"UsedLength":11.81,"UsedHeight":0}]`