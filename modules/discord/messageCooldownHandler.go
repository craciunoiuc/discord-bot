package discord

import (
	"time"

	"github.com/craciunoiuc/discord-bot/internal/types"
)

const typingStartString = "typingStart"
const presenceOnlineString = "presenceOnline"
const presenceOfflineString = "presenceOffline"

type MessageCooldownHandler struct {
	lastTimeMap map[types.Pair]time.Time
	cooldowns   map[string]time.Duration
}

var messageCooldownHandler MessageCooldownHandler

func init() {
	messageCooldownHandler = MessageCooldownHandler{
		lastTimeMap: make(map[types.Pair]time.Time),
		cooldowns:   make(map[string]time.Duration),
	}

	messageCooldownHandler.cooldowns[typingStartString] = time.Duration(5 * time.Minute)
	messageCooldownHandler.cooldowns[presenceOnlineString] = time.Duration(5 * time.Minute)
	messageCooldownHandler.cooldowns[presenceOfflineString] = time.Duration(5 * time.Minute)
}

func (h MessageCooldownHandler) canTriggerMessage(name string, userId string) bool {
	lastTime, found := h.lastTimeMap[types.Pair{X: name, Y: userId}]
	messageCooldown := h.cooldowns[name]

	return !found || !lastTime.UTC().Add(messageCooldown).After(time.Now().UTC())
}

func (h MessageCooldownHandler) triggerMessage(name string, userId string) {
	h.lastTimeMap[types.Pair{X: name, Y: userId}] = time.Now().UTC()
}

func (h MessageCooldownHandler) CanTriggerTypingStart(userId string) bool {
	return h.canTriggerMessage(typingStartString, userId)
}

func (h MessageCooldownHandler) CanTriggerPresenceOnline(userId string) bool {
	return h.canTriggerMessage(presenceOnlineString, userId)
}

func (h MessageCooldownHandler) CanTriggerPresenceOffline(userId string) bool {
	return h.canTriggerMessage(presenceOfflineString, userId)
}

func (h MessageCooldownHandler) TriggerTypingStart(userId string) {
	h.triggerMessage(typingStartString, userId)
}

func (h MessageCooldownHandler) TriggerPresenceOnline(userId string) {
	h.triggerMessage(presenceOnlineString, userId)
}

func (h MessageCooldownHandler) TriggerPresenceOffline(userId string) {
	h.triggerMessage(presenceOfflineString, userId)
}
