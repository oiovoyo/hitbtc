package hitbtc

import (
	"fmt"
	"testing"
	"time"
)

func TestHitbtc_GetMarkets(t *testing.T) {
	n := NewHitbtc(ACCESS, SECRET)

	func() {
		t := time.Now()
		defer func() {
			fmt.Printf("GetMarkets() cost %.2fs\n", time.Since(t).Seconds())
		}()
		fmt.Println(n.GetMarkets())
		//fmt.Printf("%+v\n", m)
	}()

}

func TestHitbtc_GetOrderBook(t *testing.T) {
	n := NewHitbtc(ACCESS, SECRET)
	func() {
		t := time.Now()
		defer func() {
			fmt.Printf("GetOrderBook(\"LTCBTC\") cost %.2fs\n", time.Since(t).Seconds())
		}()
		m, err := n.GetOrderBook("LTCBTC")

		fmt.Println(err)
		a := m.Buy
		for i := 0; i < len(a)-1; i++ {
			if a[i].Rate < a[i+1].Rate {
				fmt.Printf("%d %+v %+v error\n", i, a[i], a[i+1])
				break
			}
		}
		fmt.Println(a[:5])
		a = m.Sell
		for i := 0; i < len(a)-1; i++ {
			if a[i].Rate > a[i+1].Rate {
				fmt.Printf("%d %+v %+v error\n", i, a[i], a[i+1])
				break
			}
		}
		fmt.Println(a[:5])
		//fmt.Printf("%+v\n", m)
		//fmt.Printf("%+v\n", m)
	}()

}

func TestHitbtc_GetBalance(t *testing.T) {
	n := NewHitbtc(ACCESS, SECRET)
	func() {
		t := time.Now()
		defer func() {
			fmt.Printf("GetTradeBalance() cost %.2fs\n", time.Since(t).Seconds())
		}()
		tb, err := n.GetTradeBalance()

		fmt.Println(err)
		fmt.Printf("%+v\n", tb)
		pb, err := n.GetPaymentBalance()

		fmt.Println(err)
		fmt.Printf("%+v\n", pb)
	}()
}

func TestHitbtc_GetDepositAddress(t *testing.T) {
	n := NewHitbtc(ACCESS, SECRET)
	func() {
		t := time.Now()
		defer func() {
			fmt.Printf("GetDepositAddress() cost %.2fs\n", time.Since(t).Seconds())
		}()

		mapCoinAddress := map[string]string{
			"1ST":    "0x6c4fcfc5654c65517138ef636e50e68a969ce7d0",
			"8BT":    "",
			"ADX":    "0x4d6b206d0b8ed3786016ede92be2e4f6fd7df061",
			"AE":     "0x5f084418df6d24f03fa158e08962b1fe624f74c5",
			"AEON":   "WmsDVoBsc14Xjb4Hea5hHdPsC85hWtAaFFPDKs1yJTPNi2rT78mzv97hFaaEttMUt9c61Em6dP1WeHkyDtyRgWf11Q6LaKRxQ",
			"AIR":    "",
			"AMB":    "",
			"AMP":    "14Fsse7voefPJMxVjovnjPS2DmuhnGcffE",
			"ANT":    "0x97a0ae87b954ef94d305ca0b5fbf1e2581c33c77",
			"ARDR":   "360491698375",
			"AVT":    "0x1c115c336753992ad2907a5f1c196886b6d78382",
			"BAS":    "0x126213fdce80cd4949d282d3070f2c0c86dae83e",
			"BCC":    "1Mm7zwCcjRxjZV4nCJUW8detwuKJ2U8gCX",
			"BCN":    "28uUBCBMSJyYDxcoUJm8ePFtkhbtcQuxr8XaTRYdabv9JfyQV3HPHXLdi7ok6B5SQT6UXUtQgusruCoXbqUZm8VJAhhYrjk",
			"BET":    "0xe7da74daafc6bfcf1b519ec89343df97b4669b1a",
			"BMC":    "0x8103dc36cfae8055126b601d913226d411e7a70e",
			"BNT":    "0x80afcc1ab2f1c803a3aea8c84f2f192eeaf8cb5a",
			"BOS":    "",
			"BQX":    "0xe37ed72bf0e0c5066cb8c922061221610365193b",
			"BTC":    "1FWuVg9HdArjpJeWE4LFWrGVdJGZUbj793",
			"BUS":    "0xde4631786e1b6f252b9efacced55f7a33967e98e",
			"CAT":    "",
			"CDT":    "0x3b782c249586b26b538a7bb1f267ceacaa42802d",
			"CFI":    "0x6661b9537ed6ec668628243be9e7392517351a12",
			"CND":    "",
			"COSS":   "",
			"CRS":    "",
			"CSNO":   "",
			"CVC":    "0x04b68fa51d9432af1af2e8f446038ac13f7babae",
			"DASH":   "XoRwoRCpKgTbLryiLM8BEo4PhchwPQj6FS",
			"DCT":    "",
			"DDF":    "0x3b3b44e5d6e647684129b4a31fe4283240548fb7",
			"DENT":   "0x226298ad9cbc099879fb4516e6a5cf59bf06836d",
			"DGD":    "0x100901741d5c2fff3d05d204fec6e0577610bdd7",
			"DICE":   "0x9b48ffddd50d6c74b1f868986d4ad7a711c8200b",
			"DLT":    "",
			"DNT":    "0xf6d9cd98bed96b420f54c17afc1cd2f4bb398186",
			"DOGE":   "DSoDoKyBrS5mcvumqyWNcZJB3CkXf7FfHg",
			"DSH":    "DBGWqpbJNWnNDXzLY3CUSW8BgWZVLex2Ne5hhmM8KmvgSWFyf8CUFL8V3t9ZzAVrKd1mzeWaYMgFTiUHmsmiDcNt2rNw792",
			"ECAT":   "",
			"EDG":    "0xe33a390952138c61ab622d54be7fd6c5fac4b679",
			"EMC":    "EavhmzMvw5xaP5pdmrRS73WLKoS5vUttr5",
			"EMGO":   "0xa1d1a10766af21db9d612f0775d8be1df21f2253",
			"EOS":    "0x00d5066cf1dedf49e660ab25b7a8b6bec3a38073",
			"ETC":    "0x97ad07b54f2a240954b928a26a02bcd74774b976",
			"ETH":    "0x1d31570f399469ff0f0d1aeb2d7819b6207dcdd3",
			"EUR":    "",
			"EVX":    "",
			"FCN":    "6jQB6DFGKq94cAivQ91fbeAo9T87NTHFB5At3RUE7jcUaPHxtCmXcerDWZgC9gig96dF3c2qLLwD8AT8Q6xV13UiHw3Visx",
			"FUEL":   "",
			"FUN":    "0x2a96750f2b97c6d87ceef33388be7be4f5b4ade6",
			"FYN":    "0xcd39b9b7fb5aeaae7275a978a8ab03fa08e97849",
			"FYP":    "",
			"GAME":   "",
			"GNO":    "0xbd05a919da4fae014527b9fbbe3cdeaceb76d052",
			"GRPH":   "",
			"GUP":    "0x71b33ebe8797131077d246bf3460fb3a0def6764",
			"HRB":    "",
			"HVN":    "0x94fc6b609b9259dc1cfd96f5329d1a2bfe6248a2",
			"ICN":    "0xb9ddad87bbac5c90a56107574265066352fbf98d",
			"ICO":    "0x2c61c3f2e6e76eecaad61ee6ed3f569d2fb12594",
			"ICOS":   "",
			"ICX":    "",
			"IGNIS":  "",
			"IML":    "",
			"IND":    "",
			"IXT":    "0x15ad111cd3d3f5dc469151f43bfde9172134b977",
			"LRC":    "",
			"LSK":    "",
			"LTC":    "LaaQ6hDgQD1XeBdRP8oQGasf5YwvFib6wv",
			"LUN":    "0x8d4796447503b312aa403df87342c9312cf76473",
			"MAID":   "17kwddVFpNqeoqau5UtgUNQcXRGPHKbNRV",
			"MANA":   "0x55fa8ab077350a2cb8b8cc73e282a02c3468f84f",
			"MCAP":   "0xaa63de4cda39692bfcc0d349a998fb0ddf29ec86",
			"MIPS":   "",
			"MNE":    "0x95876fca2b192904d0a5e13e50ba8add45ce9e51",
			"MPK":    "",
			"MRV":    "",
			"MSP":    "0xade0bd5796b8e71d75031f4c0fec8bdb30fddaf8",
			"MTH":    "",
			"MYB":    "0x78d38cfd04f4a54a23ddb3a51a9907d5c5975296",
			"NDC":    "0x7753d08de30e9829b30b0ec3874ba6db7c36a3ff",
			"NEO":    "",
			"NET":    "0x0dd769c494e46b720dbd61f6bd33870e3e4e3498",
			"NTO":    "0x9f0a2aee2c97520bd7d6365ebfd289b4d7b8e932",
			"NXC":    "0xd5b4d97cc1e19e989b645086009c8789086408c2",
			"NXT":    "",
			"OAX":    "0x19b5a51be7e97d418258dfd93aa100874d75d4dd",
			"OMG":    "0x7125113000951943a3dd053925774f14793b4a10",
			"OPT":    "0x0fb3c75b2629263d37320d34cb6efa06aff65cd1",
			"ORME":   "",
			"PAY":    "0x2b088c808c8cdd2d704181032de9f956d98051b8",
			"PBKX":   "",
			"PING":   "",
			"PIX":    "",
			"PLBT":   "0x1c392bec84b0660488a9cefe8260f7714e7f7f18",
			"PLR":    "0xac2477d353434b07e378576735889409b2a8742c",
			"PLU":    "0xd7493eae7d262e49f542491ecf9949f4c53cd56d",
			"POE":    "0xb2b488f9b623b42efc0f5db241dc89b8f053b5b6",
			"PPC":    "PSggXaRWRDVR5QzSJ4FJVcTGfjhsRzjtes",
			"PPT":    "",
			"PQT":    "",
			"PRG":    "",
			"PRO":    "0x5d91f9ccfccd6c9ffe74d6be0bf2c8a11a7e18a8",
			"PTOY":   "0x8b32b0cab495eb0b66dcb0914c20865c3eddf3ea",
			"QAU":    "0x39b1d5bba2a87866f80845517fd1fa8f27e65409",
			"QCN":    "1Qh63ewkoCXaVxcyjDXm9aT2yqY745gLybb4zaG3LWJdCSXCAjKdyNtL1XfVjXGVLE3HP6uvQdnrLf4X4KCmKJwf1Gjhi9y",
			"QTUM":   "0xec18384f835a65039339a84c689889501865c98d",
			"REP":    "0x42a4dbb402c5610f75996d728f27d4ffe84bb8ff",
			"RKC":    "",
			"RLC":    "0x6f3e0aabd7d5d90bd04c4ca2c12e260b7d11d3b3",
			"ROOTS":  "",
			"RVT":    "0x8017f4b6ed18c878231b21fa859c2a4f0bc8f3a9",
			"SAN":    "0xf96f70c886273e8c255a5dcce7b9abff33173626",
			"SBD":    "",
			"SC":     "dd0342d468ffdd24e34d9e634288a8436caf57877606a71251c579ddaebe19ddaa483ef07449",
			"SKIN":   "0x49d7ba4ae9cb33237de285ec55b53f196e404d0d",
			"SNC":    "0xee4bc72e388b74e328bc0dba327773d40c830303",
			"SNGLS":  "0xf8884efffe933bde898a08f0ae598b0e9cd11cd5",
			"SNM":    "0x9c1bc7336b97883fd57793804865da4a9850a09c",
			"SNT":    "0x62c390ac6b0319c1ea1299e82b0dc1cefc8b9fa0",
			"STEEM":  "",
			"STRAT":  "SVzzTsHeE8dajGfnvtPnYeTcbKxV2YTh49",
			"STX":    "0x5b30ff519b06e37788e2a03d8f3793308250110a",
			"SUR":    "0x60c27a0ebf8d7c29353c84103d2b8e3b89cf0d36",
			"SWT":    "0x84a7228fe4130c103b01f3ba0c39a59b6151cb59",
			"TAAS":   "0x2f372a8b3037e60c30e2ab4fc4efcc456e86f5b9",
			"TFL":    "0x06e577fe22f905e6bc4d553ef72ee1c4b19d51ee",
			"TIME":   "0x6be5d074dd3222839543f3d9be469a4038992bd8",
			"TIX":    "0xe49e35e80ace347ddf42e81ac9981e0ee3a99c82",
			"TKN":    "0xf6a0e7ca3cacee2eb16a928b420a1633ed50f9f9",
			"TKR":    "",
			"TNT":    "0x4d77e1fd65dba05549e42d439d5ed2fbf6a324ac",
			"TRST":   "0x481b17955b578447b5e05e8ba7a10e03f0e0bcf9",
			"UET":    "0x6a72a2a45583507cdbe360ca158fc88216123f3e",
			"USD":    "12TLd269ks73rr7JFGEHDV54vt7qtDYyPo",
			"VERI":   "0x6aa90e407ec6be9bcc5f533f52cd4bc940b600be",
			"WAVES":  "3PKjxzEcUj6WiHc51xCWuqNdd61nmhWg9pd",
			"WEALTH": "0x1bd1285d859f282a07d1e4b1e77668244144f658",
			"WINGS":  "0x4a2bc492166d77317a5c7f4c2dde8c13df1a396d",
			"WMGO":   "",
			"WTT":    "",
			"XAUR":   "0x2b3cc7e261e63b34e3fc1a4b81911e5d98ba57a9",
			"XDN":    "ddd22hKTzZfaPRoKYjFgs48hgoqq63HHSjDjQ2eV7huvbjBQypZcwnCgUw7wMnJH82FS8noajsHKGiNABUygXMdm1oKkiLNPF",
			"XDNCO":  "",
			"XEM":    "",
			"XLC":    "",
			"XMR":    "49WFQmEb165J59Kgrh6PG51xnG5Zg3JxyizagCfKbn2fMLXAGEGgvDxjXqrT3anyZ22j7DEE74GkbVcQFyH2nNiC3dv2QYD",
			"XRP":    "",
			"XTZ":    "",
			"YOYOW":  "",
			"ZEC":    "t1X27YWHgC4Rkyv7GQVHcWCgEeq19vqGjqq",
			"ZRC":    "",
			"ZRX":    "0x25b899a0517cc40e83e7d43733646642d6da99b1",
		}

		for k, _ := range mapCoinAddress {

			if mapCoinAddress[k] == "" {
				continue
			}

			func() {
				t1 := time.Now()
				defer func() {
					fmt.Printf("GetDepositAddress(\"%s\") cost %.2fs\n", k, time.Since(t1).Seconds())
				}()
				for i := 0; i < 2; i++ {
					addr, err := n.GetDepositAddress(k)
					if err != nil {
						fmt.Printf("get %s address error %v\n", k, err)
						time.Sleep(time.Millisecond * 500)
						continue

					}
					//fmt.Printf("get address %s %s\n", k, addr)
					if mapCoinAddress[k] != addr {
						fmt.Printf("wtf %s %s != %s\n", k, mapCoinAddress[k], addr)
					}
					break
				}
				time.Sleep(time.Millisecond * 100)
			}()
			//a, _ := json.Marshal(mapCoinAddress)
			//fmt.Println(string(a))

		}

	}()
}
func TestHitbtc_GetDepositAddressID(t *testing.T) {
	n := NewHitbtc(ACCESS, SECRET)
	func() {
		t := time.Now()
		defer func() {
			fmt.Printf("GetDepositAddress() cost %.2fs\n", time.Since(t).Seconds())
		}()
		//fmt.Println(n.GetDepositAddress("BTC"))
		// to be solved
		fmt.Println(n.GetDepositAddress("NXT"))
	}()
}

func TestHitbtc_Withdraw(t *testing.T) {
	n := NewHitbtc(ACCESS, SECRET)
	func() {
		t := time.Now()
		defer func() {
			fmt.Printf("Withdraw() cost %.2fs\n", time.Since(t).Seconds())
		}()
		//fmt.Println(n.GetDepositAddress("BTC"))
		// to be solved
		fmt.Println(n.Withdraw("BTC", 0.09985508, "1KyjuCs3MPeD7SS8YaUTbrwumeNfEqL533", ""))
	}()
}

func TestHitbtc_TransferToMain(t *testing.T) {
	n := NewHitbtc(ACCESS, SECRET)
	func() {
		t := time.Now()
		defer func() {
			fmt.Printf("Withdraw() cost %.2fs\n", time.Since(t).Seconds())
		}()
		//fmt.Println(n.GetDepositAddress("BTC"))
		// to be solved
		fmt.Println(n.TransferToMain("BTC", 0.09985508))
	}()
}

func TestHitbtc_TransferToTrading(t *testing.T) {
	n := NewHitbtc(ACCESS, SECRET)
	func() {
		t := time.Now()
		defer func() {
			fmt.Printf("Withdraw() cost %.2fs\n", time.Since(t).Seconds())
		}()
		//fmt.Println(n.GetDepositAddress("BTC"))
		// to be solved
		fmt.Println(n.TransferToTrading("BTC", 0.1))
	}()
}

func TestHitbtc_adjustPriceAmount(t *testing.T) {
	n := NewHitbtc(ACCESS, SECRET)
	func() {
		t := time.Now()
		defer func() {
			fmt.Printf("Withdraw() cost %.2fs\n", time.Since(t).Seconds())
		}()
		//fmt.Println(n.GetDepositAddress("BTC"))
		// to be solved
		fmt.Println(n.adjustPriceAmount("ETHBTC", 0.05, 0.0222221))
	}()
}

func TestHitbtc_BuyLimit(t *testing.T) {
	n := NewHitbtc(ACCESS, SECRET)
	func() {
		t := time.Now()
		defer func() {
			fmt.Printf("Withdraw() cost %.2fs\n", time.Since(t).Seconds())
		}()
		//fmt.Println(n.GetDepositAddress("BTC"))
		// to be solved
		fmt.Println(n.BuyLimit("ETHBTC", 0.5, 0.06845555))
	}()
}

func TestHitbtc_SellLimit(t *testing.T) {
	n := NewHitbtc(ACCESS, SECRET)
	func() {
		t := time.Now()
		defer func() {
			fmt.Printf("Withdraw() cost %.2fs\n", time.Since(t).Seconds())
		}()
		//fmt.Println(n.GetDepositAddress("BTC"))
		// to be solved
		fmt.Println(n.SellLimit("ETHBTC", 0.5, 0.06815555555))
	}()
}

func TestHitbtc_CancelOrder(t *testing.T) {
	n := NewHitbtc(ACCESS, SECRET)
	func() {
		t := time.Now()
		defer func() {
			fmt.Printf("CancelOrder() cost %.2fs\n", time.Since(t).Seconds())
		}()
		//fmt.Println(n.GetDepositAddress("BTC"))
		// to be solved
		fmt.Println(n.CancelOrder("iYymCZgUhmHYqWWicyXXjoCNpxLBosej"))
	}()
}

func TestHitbtc_CancelOrders(t *testing.T) {
	n := NewHitbtc(ACCESS, SECRET)
	func() {
		t := time.Now()
		defer func() {
			fmt.Printf("CancelOrders() cost %.2fs\n", time.Since(t).Seconds())
		}()
		//fmt.Println(n.GetDepositAddress("BTC"))
		// to be solved
		fmt.Println(n.CancelOrders("ETHBTC"))
	}()
}

func TestHitbtc_GetDepositWithdrawal(t *testing.T) {
	n := NewHitbtc(ACCESS, SECRET)
	func() {
		t := time.Now()
		defer func() {
			fmt.Printf("GetDepositWithdrawal() cost %.2fs\n", time.Since(t).Seconds())
		}()
		//fmt.Println(n.GetDepositAddress("BTC"))
		// to be solved
		fmt.Println(n.GetDepositWithdrawal("BTC", 0, 100))
	}()
}

func TestHitbtc_GetOrder(t *testing.T) {
	n := NewHitbtc(ACCESS, SECRET)
	func() {
		t := time.Now()
		defer func() {
			fmt.Printf("GetOrder() cost %.2fs\n", time.Since(t).Seconds())
		}()
		//fmt.Println(n.GetDepositAddress("BTC"))
		// to be solved
		fmt.Println(n.GetOrder("iYymCZgUhmHYqWWicyXXjoCNpxLBosej"))
	}()
}
func TestHitbtc_GetPaymentBalance(t *testing.T) {
	n := NewHitbtc(ACCESS, SECRET)
	func() {
		t := time.Now()
		defer func() {
			fmt.Printf("GetPaymentBalance() cost %.2fs\n", time.Since(t).Seconds())
		}()
		//fmt.Println(n.GetDepositAddress("BTC"))
		// to be solved
		fmt.Println(n.GetPaymentBalance())
	}()
}
func TestHitbtc_GetTradeBalance(t *testing.T) {
	n := NewHitbtc(ACCESS, SECRET)
	func() {
		t := time.Now()
		defer func() {
			fmt.Printf("GetTradeBalance() cost %.2fs\n", time.Since(t).Seconds())
		}()
		//fmt.Println(n.GetDepositAddress("BTC"))
		// to be solved
		fmt.Println(n.GetTradeBalance())
	}()
}

func TestHitbtc_GetTransaction(t *testing.T) {
	n := NewHitbtc(ACCESS, SECRET)
	func() {
		t := time.Now()
		defer func() {
			fmt.Printf("GetTransaction() cost %.2fs\n", time.Since(t).Seconds())
		}()
		//fmt.Println(n.GetDepositAddress("BTC"))
		// to be solved
		fmt.Println(n.GetTransaction("846cf968-0ec5-45b1-aaff-b472597a1b54"))
	}()
}

const (
	ACCESS = ""
	SECRET = ""
)
