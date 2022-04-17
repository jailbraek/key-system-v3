package main

import (
	"DarkHub-KeySys-V3/utils"
	"encoding/json"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/monitor"
	"log"
	"os"
	"time"
)

const (
	HOME           = "b2"
	CHECKPOINT1    = "3g"
	CHECKPOINT2    = "BH"
	KEY            = "Ve"
	BID            = "eb"
	STAFFK         = "sfd"
	VERSION        = "v3.0"
	IDENTITY       = "DARKHUB-v3.0-WEUGo5YJPkFe4eMH5V2Cvkq3oHFYuV"
	Checkpoint1Url = "https://work.ink/l/1n8/DarkHubKey1"
	Checkpoint2Url = "https://work.ink/en/l/1n8/DarkHubKey2"
	keyFiller      = "penispenispenispenispenispenispenispenispenispenispenispenispenispenispenispenispenispenispenispenispenispenispenis"
)

var (
	key        = []byte("iHOFtYu6Hv0kQz6%ZMf2G1!VM76aD2f!")
	keyGenKey  = []byte("8If05g51m6uF&Oe#0QZGUb4#j2rKVizb")
	keyStubKey = []byte("8If05g51m6uF&Oe#0QZGUb4#j2rKVizb")
)

/*
	enc, err := utils.Encrypt([]byte("ADAM LIKE MEN"), key)
	if err != nil {
		panic(err)
	}
	fmt.Println(enc)
	normal, err := utils.Decrypt(enc, key)
	if err != nil {
		panic(err)
	}
	fmt.Println(normal)
*/

func main() {
	app := fiber.New(fiber.Config{
		Prefork:           true,
		ServerHeader:      "Your Mom's fat cock",
		ProxyHeader:       "CF-Connecting-IP",
		AppName:           "Dark-Key Sys v3",
		ReduceMemoryUsage: true,
	})
	app.Use("/discord", func(c *fiber.Ctx) error {
		data, err := os.ReadFile("discordinvite.txt")
		if err != nil {
			return c.SendStatus(500)
		}
		return c.Redirect(string(data))
	})
	app.Use("/41BK2NJz9Vond7rYrbAF", monitor.New())
	app.Get("/", func(c *fiber.Ctx) error {
		ip := utils.HashIP(c.IP())
		bid := c.Cookies(BID)
		k := c.Cookies(KEY)
		if len(k) != 0 {
			k, err := utils.Decrypt(k, key)
			if err != nil {
				c.ClearCookie(KEY)
				return c.Redirect("/")
			}
			if string(k) != keyFiller {
				ck, err := utils.CheckKey(k, ip, VERSION, keyGenKey)
				if err != nil {
					c.ClearCookie(KEY)
				}
				if ck {
					return c.Redirect("/I/made/this/in/3/hrs/and/it/works")
				}
			}
		}
		if len(bid) == 0 {
			bid = utils.GenerateBrowserID()
			enc, err := utils.Encrypt([]byte(bid), key)
			if err != nil {
				fmt.Println("Some kid is fucking with bid: ", err)
				c.Set("refresh", "5;url=/")
				return c.SendStatus(400)
			}
			b := fiber.Cookie{
				Name:    BID,
				Value:   enc,
				Expires: time.Now().Add(time.Hour * 24 * 365 * 10),
			}
			c.Cookie(&b)
		} else {
			var err error
			bid, err = utils.Decrypt(bid, key)
			if err != nil {
				return c.Redirect("/")
			}
		}
		cp := check{
			Time:        time.Now().Unix(),
			Ip:          ip,
			BrowserID:   bid,
			CheckPoint:  0,
			DarkhubBest: true,
			Darkhub:     "darkhubdarkhubdarkhub is the best",
			Penis:       "penis penis penis penis penis penis, adam likes big men",
			AdamLikeMen: true,
		}
		cpData, err := json.Marshal(cp)
		if err != nil {
			fmt.Println("Some kid just fucked with marshalling: ", err)
			return c.Redirect(c.OriginalURL())
		}
		enc, err := utils.Encrypt(cpData, key)
		if err != nil {
			fmt.Println("Some kid just fucked with enc: ", err)
			return c.Redirect(c.OriginalURL())
		}
		cp1Data, err := utils.Encrypt([]byte(Checkpoint1Url), key)
		if err != nil {
			fmt.Println("Some kid just fucked with enc: ", err)
			return c.Redirect(c.OriginalURL())
		}
		cp1 := fiber.Cookie{
			Name:    CHECKPOINT1,
			Value:   cp1Data,
			Expires: time.Now().Add(time.Hour * 24),
		}
		c.Cookie(&cp1)
		cp2Data, err := utils.Encrypt([]byte(Checkpoint2Url), key)
		if err != nil {
			fmt.Println("Some kid just fucked with enc: ", err)
			return c.Redirect(c.OriginalURL())
		}
		cp2 := fiber.Cookie{
			Name:    CHECKPOINT2,
			Value:   cp2Data,
			Expires: time.Now().Add(time.Hour * 24),
		}
		c.Cookie(&cp2)
		keyData, err := utils.Encrypt([]byte(keyFiller), key)
		if err != nil {
			fmt.Println("Some kid just fucked with enc: ", err)
			return c.Redirect(c.OriginalURL())
		}
		key := fiber.Cookie{
			Name:  KEY,
			Value: keyData,
		}
		c.Cookie(&key)
		cookie := fiber.Cookie{
			Name:    HOME,
			Value:   enc,
			Expires: time.Now().Add(time.Hour * 24),
		}
		c.Cookie(&cookie)
		return c.Redirect(Checkpoint1Url)
	})
	app.Get("/staff/generateDonatorKey/ForMePls", func(c *fiber.Ctx) error {
		if c.Cookies(STAFFK) != "true" {
			return c.Status(404).SendString("Cannot GET /staff/generateDonatorKey/")
		}
		stub, err := utils.GenerateDonatorKeyStub(keyStubKey)
		if err != nil {
			return c.SendStatus(400)
		}
		return c.SendString(stub)
	})
	app.Get("/donator/redeem/:key", func(c *fiber.Ctx) error {
		d := c.Params("key")
		e := utils.RedeemKeyStub(d, utils.HashIP(c.IP()), VERSION, keyStubKey, keyGenKey)
		if len(e) == 0 {
			return c.Status(404).SendString("Cannot GET /donator/redeem/" + d)
		}
		fmt.Println("Donator Key Redeemed by:", utils.HashIP(c.IP()))
		enc, err := utils.Encrypt([]byte(e), key)
		if err != nil {
			return c.Status(500).SendString("Error generating key, contact a developer.")
		}
		cookie := fiber.Cookie{
			Name:    KEY,
			Value:   enc,
			Expires: time.Now().Add(time.Hour * 24 * 365 * 10),
		}
		c.Cookie(&cookie)
		return c.SendString(e)
	})
	checkpoints := app.Group("/checkpoints")
	checkpoints.Get("/1", func(c *fiber.Ctx) error {
		if c.GetReqHeaders()["Referer"] != "https://work.ink/" {
			return c.Redirect("/")
		}
		ip := utils.HashIP(c.IP())
		home := c.Cookies(HOME)
		if len(home) == 0 {
			return c.Redirect("/")
		}
		dec, err := utils.Decrypt(home, key)
		if err != nil {
			fmt.Println("checkpoint 1 dec err: ", err)
			return c.Redirect("/")
		}
		var cp check
		err = json.Unmarshal([]byte(dec), &cp)
		if err != nil {
			fmt.Println("checkpoint 1 unmarshal err: ", err)
			fmt.Println(dec)
			return c.Redirect("/")
		}
		bid := c.Cookies(BID)
		if len(bid) == 0 {
			return c.Redirect("/")
		}
		bid, err = utils.Decrypt(bid, key)
		if err != nil {
			fmt.Println("checkpoint 1 bid dec err: ", err)
			return c.Redirect("/")
		}
		if cp.Ip != ip || bid != cp.BrowserID || cp.Time < time.Now().Unix()-((60*10)*1000) {
			c.ClearCookie(HOME)
			return c.Redirect("/")
		}

		check := check{
			Ip:          ip,
			Time:        time.Now().Unix(),
			BrowserID:   bid,
			CheckPoint:  1,
			DarkhubBest: true,
			Darkhub:     "darkhubdarkhubdarkhub is the best",
			Penis:       "penis penis penis penis penis penis, adam likes big men",
			AdamLikeMen: true,
		}
		data, err := json.Marshal(check)
		if err != nil {
			fmt.Println("checkpoint 1 dec err: ", err)
			return c.Redirect("/")
		}
		enc, err := utils.Encrypt(data, key)
		if err != nil {
			fmt.Println("Some kid just fucked with enc: ", err)
			return c.Redirect(c.OriginalURL())
		}
		cookie := fiber.Cookie{
			Name:    CHECKPOINT1,
			Value:   enc,
			Expires: time.Now().Add(time.Hour * 24),
		}
		c.Cookie(&cookie)
		return c.Redirect(Checkpoint2Url)
	})
	checkpoints.Get("/2", func(c *fiber.Ctx) error {
		if c.GetReqHeaders()["Referer"] != "https://work.ink/" {
			return c.Redirect("/")
		}
		ip := utils.HashIP(c.IP())
		home := c.Cookies(HOME)
		cp1Data := c.Cookies(CHECKPOINT1)
		if len(home) == 0 || len(cp1Data) == 0 {
			return c.Redirect("/")
		}
		dec2, err := utils.Decrypt(cp1Data, key)
		dec, err := utils.Decrypt(home, key)
		if err != nil {
			fmt.Println("checkpoint 1 dec err: ", err)
			return c.Redirect("/")
		}
		var cp check
		var cp1 check
		err = json.Unmarshal([]byte(dec), &cp)
		if err != nil {
			fmt.Println("checkpoint 1 unmarshal err: ", err)
			fmt.Println(dec)
			return c.Redirect("/")
		}
		err = json.Unmarshal([]byte(dec2), &cp1)
		if err != nil {
			fmt.Println("checkpoint 1 unmarshal err: ", err)
			fmt.Println(dec2)
			return c.Redirect("/checkpoints/1")
		}
		bid := c.Cookies(BID)
		if len(bid) == 0 {
			return c.Redirect("/")
		}
		bid, err = utils.Decrypt(bid, key)
		if err != nil {
			fmt.Println("checkpoint 1 bid dec err: ", err)
			return c.Redirect("/")
		}
		if cp.Ip != ip || cp1.Ip != ip {
			c.ClearCookie(HOME)
			c.ClearCookie(CHECKPOINT1)
			return c.Redirect("/")
		}
		if bid != cp.BrowserID || bid != cp1.BrowserID {
			c.ClearCookie(HOME)
			c.ClearCookie(CHECKPOINT1)
			return c.Redirect("/")
		}
		if cp.Time < time.Now().Unix()-((60*14)*1000) || cp1.Time < time.Now().Unix()-((60*12)*1000) {
			c.ClearCookie(HOME)
			c.ClearCookie(CHECKPOINT1)
			return c.Redirect("/")
		}
		check := check{
			Ip:          ip,
			Time:        time.Now().Unix(),
			BrowserID:   bid,
			CheckPoint:  2,
			DarkhubBest: true,
			Darkhub:     "darkhubdarkhubdarkhub is the best",
			Penis:       "penis penis penis penis penis penis, adam likes big men",
			AdamLikeMen: true,
		}
		data, err := json.Marshal(check)
		if err != nil {
			fmt.Println("checkpoint 1 dec err: ", err)
			return c.Redirect("/")
		}
		enc, err := utils.Encrypt(data, key)
		if err != nil {
			fmt.Println("Some kid just fucked with enc: ", err)
			return c.Redirect(c.OriginalURL())
		}
		cookie := fiber.Cookie{
			Name:    CHECKPOINT2,
			Value:   enc,
			Expires: time.Now().Add(time.Hour * 24),
		}
		c.Cookie(&cookie)
		return c.Redirect("/Adam/Like/Big/Men/DarkhubOnTop/Pe/n/i/s/GetKey")
	})
	app.Get("/Adam/Like/Big/Men/DarkhubOnTop/Pe/n/i/s/GetKey", func(c *fiber.Ctx) error {
		ip := utils.HashIP(c.IP())
		homeData := c.Cookies(HOME)
		cp1Data := c.Cookies(CHECKPOINT1)
		cp2Data := c.Cookies(CHECKPOINT2)
		bid := c.Cookies(BID)
		if len(homeData) == 0 || len(cp1Data) == 0 || len(cp2Data) == 0 || len(bid) == 0 {
			c.ClearCookie(HOME)
			c.ClearCookie(CHECKPOINT1)
			c.ClearCookie(CHECKPOINT2)
			return c.Redirect("/")
		}
		bid, err := utils.Decrypt(bid, key)
		if err != nil {
			fmt.Println("bid dec err: ", err)
			return c.Redirect("/")
		}
		homeData, err = utils.Decrypt(homeData, key)
		if err != nil {
			fmt.Println("home dec err: ", err)
			return c.Redirect("/")
		}
		cp1Data, err = utils.Decrypt(cp1Data, key)
		if err != nil {
			fmt.Println("checkpoint 1 dec err: ", err)
			return c.Redirect("/")
		}
		cp2Data, err = utils.Decrypt(cp2Data, key)
		if err != nil {
			fmt.Println("checkpoint 2 dec err: ", err)
			return c.Redirect("/")
		}
		var home check
		var cp1 check
		var cp2 check
		err = json.Unmarshal([]byte(homeData), &home)
		if err != nil {
			fmt.Println("home unmarshal err: ", err)
			fmt.Println(homeData)
			return c.Redirect("/")
		}
		err = json.Unmarshal([]byte(cp1Data), &cp1)
		if err != nil {
			fmt.Println("checkpoint 1 unmarshal err: ", err)
			fmt.Println(cp1Data)
			return c.Redirect("/")
		}
		err = json.Unmarshal([]byte(cp2Data), &cp2)
		if err != nil {
			fmt.Println("checkpoint 2 unmarshal err: ", err)
			fmt.Println(cp2Data)
			return c.Redirect("/")
		}
		if home.Ip != ip || cp1.Ip != ip || cp2.Ip != ip || home.BrowserID != bid || cp1.BrowserID != bid || cp2.BrowserID != bid || home.Time < time.Now().Unix()-((60*16)*1000) || cp1.Time < time.Now().Unix()-((60*14)*1000) || cp2.Time < time.Now().Unix()-((60*12)*1000) {
			c.ClearCookie(HOME)
			c.ClearCookie(CHECKPOINT1)
			c.ClearCookie(CHECKPOINT2)
			return c.Redirect("/")
		}
		k, err := utils.GenerateKey(ip, VERSION, keyGenKey, false, 0, time.Now().Add(time.Hour*24).Unix())
		if err != nil {
			fmt.Println("key gen err: ", err)
			return c.Redirect("/")
		}
		if b, err := utils.CheckKey(k, ip, VERSION, keyGenKey); !b || err != nil {
			k, err = utils.GenerateKey(ip, VERSION, keyGenKey, false, 0, time.Now().Add(time.Hour*24).Unix())
		}
		enc, err := utils.Encrypt([]byte(k), key)
		if err != nil {
			fmt.Println("key enc err: ", err)
			return c.Redirect(c.OriginalURL())
		}
		cookie := fiber.Cookie{
			Name:    KEY,
			Value:   enc,
			Expires: time.Now().Add(time.Hour * 24),
		}
		c.Cookie(&cookie)
		return c.SendString(k)
	})
	app.Get("/I/made/this/in/3/hrs/and/it/works", func(c *fiber.Ctx) error {
		k := c.Cookies(KEY)
		if len(k) == 0 {
			return c.Redirect("/")
		}
		k, err := utils.Decrypt(k, key)
		vaild, err := utils.CheckKey(k, utils.HashIP(c.IP()), VERSION, keyGenKey)
		if err != nil {
			fmt.Println("key check err: ", err)
			c.ClearCookie(KEY)
			return c.Redirect("/")
		}
		if !vaild {
			fmt.Println("key check failed")
			c.ClearCookie(KEY)
			return c.Redirect("/")
		}
		return c.SendString(k)
	})
	app.Post("/lol/Adam/Wants/Sonic/CheckKey/rn/pls", func(c *fiber.Ctx) error {
		ip := utils.HashIP(c.IP())
		var d checkKeyBody
		err := c.BodyParser(&d)
		if err != nil {
			fmt.Println("body parser err: ", err)
			return c.Redirect("/")
		}
		if d.Key == "" || len(d.Key) <= 100 {
			return c.SendStatus(400)
		}
		if b, err := utils.CheckKey(d.Key, ip, VERSION, keyGenKey); !b || err != nil {
			return c.SendStatus(400)
		}
		k, err := utils.ParseKey(d.Key, keyGenKey)
		if err != nil {
			fmt.Println("key parse err: ", err)
			return c.Redirect("/")
		}
		var sd = sendData{
			Key:      d.Key,
			Ip:       k.Ip,
			Version:  k.Version,
			Donator:  k.Donator,
			Identity: IDENTITY,
		}
		b, err := json.Marshal(sd)
		if err != nil {
			fmt.Println("json marshal err: ", err)
			return c.Redirect(c.OriginalURL())
		}
		return c.SendString(utils.FunnyEncoding(b))
	})
	app.Get("/what/the/he/ll/is/my/ip", func(c *fiber.Ctx) error {
		if c.GetReqHeaders()["Id"] != IDENTITY {
			return c.SendStatus(403)
		}
		ip := utils.HashIP(c.IP())
		return c.SendString(ip)
	})
	app.Get("/pingCheck", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"status": "up",
		})
	})
	log.Fatalln(app.Listen(":5001"))
}

type (
	check struct {
		Time        int64  `json:"time"`
		Ip          string `json:"ip"`
		BrowserID   string `json:"browserID"`
		CheckPoint  int    `json:"checkPoint"`
		DarkhubBest bool   `json:"darkhubBest"`
		Darkhub     string `json:"darkhub"`
		Penis       string `json:"penis"`
		AdamLikeMen bool   `json:"adamLikeMen"`
	}
	checkKeyBody struct {
		Key string `json:"key"`
	}
	sendData struct {
		Key      string `json:"key"`
		Ip       string `json:"ip"`
		Version  string `json:"version"`
		Identity string `json:"identity"`
		Donator  bool   `json:"donator"`
	}
)
