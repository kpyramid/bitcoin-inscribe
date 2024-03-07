package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/kpyramid/bitcoin-inscribe/types"
	log "github.com/sirupsen/logrus"
	"os"
	"strconv"
)

const html = `<!DOCTYPE html>
<html lang="en">
<head>
<style>
    .container img {
        position: absolute;
        width: 100%%;
        height: 100%%;
        object-fit: cover;
    }
    .container img:nth-child(1) { z-index: 1; }
    .container img:nth-child(2) { z-index: 2; }
    .container img:nth-child(3) { z-index: 3; }
    .container img:nth-child(4) { z-index: 4; }
    .container img:nth-child(5) { z-index: 5; }
    .container img:nth-child(6) { z-index: 6; }
</style>
</head>
<body>
    <div class="container">
        <img src="/content/%s">
        <img src="/content/%s">
        <img src="/content/%s">
        <img src="/content/%s">
        <img src="/content/%s">
        <img src="/content/%s">
    </div>
</body>
</html>
`

func main() {
	svc := types.GetServiceContext()

	// parse nft info
	var dataList map[int64][]string
	content, err := os.ReadFile("./recursion_data.json")
	if err != nil {
		log.Fatal(err)
	}
	if err := json.Unmarshal(content, &dataList); err != nil {
		log.Fatal(err)
	}

	for i, item := range dataList {
		if len(item) != 6 {
			log.Fatal("invalid item. %v", item)
		}
		svc.Redis.HSet(context.TODO(), types.OrderNFTDetail, strconv.FormatInt(i, 10), fmt.Sprintf(html, item[0], item[1], item[2], item[3], item[4], item[5]))
		log.Infof("set success. token_id: %d", i)
	}
}
