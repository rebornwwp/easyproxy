package structure

import "sync"

type ChannelManager struct {
	channels []Channel
	mapSrc   map[string]*Channel
	mapDst   map[string]*Channel
	mutex    *sync.Mutex
}

func (channelManager *ChannelManager) Init() {
	channelManager.channels = make([]Channel, 0)
	channelManager.mapSrc = make(map[string]*Channel)
	channelManager.mapDst = make(map[string]*Channel)
	channelManager.mutex = new(sync.Mutex)
}

func (channelManager *ChannelManager) PutChannel(channel *Channel) {
	channelManager.mutex.Lock()
	channelManager.channels = append(channelManager.channels, *channel)
	channelManager.mapSrc[channel.SrcUrl()] = channel
	channelManager.mapSrc[channel.DstUrl()] = channel
	channelManager.mutex.Unlock()
}

func (channelManager *ChannelManager) DeleteChannel(channel *Channel) {

}

func (channelManager *ChannelManager) GetChannels() []Channel {
	return channelManager.channels
}

func (channelManager *ChannelManager) Check() (error, error) {
}

func deleteMap(_map map[string]*Channel, url string) {
	_, ok := _map[url]
	if ok {
		delete(_map, url)
	}
}

func (channelManager *ChannelManager) Clean() {
	for _, channel := range channelManager.channels {
		deleteMap(channelManager.mapSrc, channel.SrcUrl())
		deleteMap(channelManager.mapDst, channel.DstUrl())
		channel.Close()
	}
	channelManager.channels = channelManager.channels[:0]
}
