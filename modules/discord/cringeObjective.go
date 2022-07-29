package discord

type CringeObjective struct {
	targetUserIds []string
}

func newCringeObjective(userIds []string) *CringeObjective {
	return &CringeObjective{
		targetUserIds: userIds,
	}
}
