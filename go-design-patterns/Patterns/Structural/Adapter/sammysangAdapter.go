package main

type sammysangAdapter struct {
	sstv *SammysangTV
}

func (ss *sammysangAdapter) turnOn() {
	ss.sstv.setOnState(true)
}

func (ss *sammysangAdapter) turnOff() {
	ss.sstv.setOnState(false)
}

func (ss *sammysangAdapter) volumeUp() int {
	vol := ss.sstv.getVolume()
	vol++
	ss.sstv.setVolume(vol)
	return vol
}

func (ss *sammysangAdapter) volumeDown() int {
	vol := ss.sstv.getVolume()
	vol--
	ss.sstv.setVolume(vol)
	return vol
}

func (ss *sammysangAdapter) channelUp() int {
	ch := ss.sstv.getChannel()
	ch++
	ss.sstv.setChannel(ch)
	return ch
}

func (ss *sammysangAdapter) channelDown() int {
	ch := ss.sstv.getChannel()
	ch--
	ss.sstv.setChannel(ch)
	return ch
}

func (ss *sammysangAdapter) goToChannel(ch int) {
	ss.sstv.setChannel(ch)
}
