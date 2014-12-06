/*
Implements methods to interact with the official Trade Offer API.

See: https://developer.valvesoftware.com/wiki/Steam_Web_API/IEconService
*/
package tradeoffer

import (
	"github.com/smithfox/gosteam/economy"
	"github.com/smithfox/gosteam/economy/inventory"
	"github.com/smithfox/gosteam/steamid"
)

type TradeOfferState uint

const (
	TradeOfferState_Invalid      TradeOfferState = 1 // Invalid
	TradeOfferState_Active                       = 2 // This trade offer has been sent, neither party has acted on it yet.
	TradeOfferState_Accepted                     = 3 // The trade offer was accepted by the recipient and items were exchanged.
	TradeOfferState_Countered                    = 4 // The recipient made a counter offer
	TradeOfferState_Expired                      = 5 // The trade offer was not accepted before the expiration date
	TradeOfferState_Canceled                     = 6 // The sender cancelled the offer
	TradeOfferState_Declined                     = 7 // The recipient declined the offer
	TradeOfferState_InvalidItems                 = 8 // Some of the items in the offer are no longer available (indicated by the missing flag in the output)
)

type Asset struct {
	AppId      uint64             `json:",string"`
	ContextId  economy.ContextId  `json:",string"`
	AssetId    economy.AssetId    `json:",string"`
	CurrencyId uint64             `json:",string"`
	ClassId    economy.ClassId    `json:",string"`
	InstanceId economy.InstanceId `json:",string"`
	Amount     uint64             `json:",string"`
	Missing    bool
}

type TradeOfferId uint64

type TradeOffer struct {
	TradeOfferId   TradeOfferId    `json:",string"`
	OtherAccountId steamid.SteamId `json:"accountid_other"`
	Message        string
	ExpirationTime uint32          `json:"expiraton_time"`
	State          TradeOfferState `json:"trade_offer_state"`
	ToGive         []*Asset        `json:"items_to_give"`
	ToReceive      []*Asset        `json:"items_to_receive"`
	IsOurs         bool            `json:"is_our_offer"`
	TimeCreated    uint32          `json:"time_created"`
	TimeUpdated    uint32          `json:"time_updated"`
}

type TradeOffers struct {
	Sent         []*TradeOffer `json:"trade_offers_sent"`
	Received     []*TradeOffer `json:"trade_offers_received"`
	Descriptions inventory.Descriptions
}

type TradeOfferSendAsset struct {
	AppId     int   `json:"appid"`
	ContextId int   `json:"contextid"`
	AssetId   int64 `json:"assetid"`
	Amount    int   `json:"amount"`
}

type TradeOfferSendAssetReady struct {
	Assets []TradeOfferSendAsset `json:"assets"`
	Ready  bool                  `json:"ready"`
}

type TradeOfferSendObj struct {
	NewVersion string                   `json:"newversion"`
	Version    int                      `json:"version"`
	Me         TradeOfferSendAssetReady `json:"me"`
	Them       TradeOfferSendAssetReady `json:"them"`
}
