package state

type TriggerResult struct {
	Trigger Trigger
	hmState HmState
}

var rules = map[HmState][]TriggerResult{
	OnHook: {
		TriggerResult{Trigger: CallDialed, hmState: Connecting},
	},
	Connecting: {
		TriggerResult{Trigger: HungUp, hmState: OnHook}, // hung up 電話を切る
		TriggerResult{Trigger: CallConnected, hmState: Connected},
	},
	Connected: {
		TriggerResult{Trigger: LeftMessage, hmState: OnHook},
		TriggerResult{Trigger: HungUp, hmState: OnHook},
		TriggerResult{Trigger: PlacedOnHold, hmState: OnHold},
	},
	OnHold: {
		TriggerResult{Trigger: TakenOffHold, hmState: Connected},
		TriggerResult{Trigger: HungUp, hmState: OnHook},
	},
}
