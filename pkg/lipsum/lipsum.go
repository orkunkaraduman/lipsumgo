package lipsum

import (
	"bytes"
	"io"
	"strings"
	"sync"
)

const Lipsum = `Lorem ipsum dolor sit amet, consectetur adipiscing elit. Proin vitae purus eu massa fringilla ultricies. Aliquam augue nisl, efficitur eu tellus eget, volutpat varius purus. Pellentesque consectetur eleifend elit vel gravida. Duis mollis sit amet tortor sit amet dignissim. Nam velit lorem, lobortis eget vestibulum et, posuere eu arcu. Integer turpis magna, mattis sed dolor vitae, pretium scelerisque nibh. Vivamus porta pellentesque metus, a congue massa egestas nec. Etiam id ligula laoreet, imperdiet augue sed, scelerisque purus. Donec cursus odio ante, eu venenatis nisl tincidunt in. Praesent vitae lorem accumsan, fringilla orci eu, vehicula odio. Aenean ac odio varius, tincidunt odio eget, tincidunt odio. In nec dolor venenatis, vulputate neque sed, faucibus neque.

Duis sed ex sit amet magna efficitur tristique in at erat. Praesent in elit magna. Donec rhoncus arcu est, non suscipit orci facilisis quis. Etiam ipsum ex, ultrices sed eros eget, molestie sagittis neque. Nam dapibus turpis at elit luctus porta. Orci varius natoque penatibus et magnis dis parturient montes, nascetur ridiculus mus. Praesent ultricies lacus eget neque volutpat tincidunt.

Suspendisse ultricies urna at dui malesuada, non elementum sapien posuere. Interdum et malesuada fames ac ante ipsum primis in faucibus. Ut ac suscipit ligula, vel molestie sapien. Quisque non tempus ante, nec interdum risus. Nam in diam eros. Etiam posuere nisi non ipsum blandit aliquam. Aenean feugiat at erat non ultrices. Proin euismod pellentesque tellus, non iaculis nulla rhoncus id. Aliquam dapibus aliquet laoreet. Nunc consequat turpis eu nulla pretium condimentum. Aliquam sed diam bibendum, facilisis lectus et, hendrerit nisi. Aliquam a elit ultrices odio commodo commodo.

Nulla quis purus ac nisi egestas rutrum a eu metus. Orci varius natoque penatibus et magnis dis parturient montes, nascetur ridiculus mus. Phasellus sed placerat nibh. Suspendisse tincidunt libero ipsum, vel eleifend mi venenatis at. Nam metus eros, sollicitudin ac faucibus sed, fermentum ac mauris. Nullam fringilla pretium velit, at fermentum enim finibus at. Proin id lacus nec nulla scelerisque tempor.

Mauris sit amet mi a ipsum porta consectetur quis vitae sem. Etiam dolor lorem, consequat vitae faucibus tempus, rutrum vel neque. Integer nulla libero, tincidunt non tempor id, dictum non sem. Etiam eget lacus nunc. Nam sed nulla vitae diam imperdiet scelerisque. In vitae sollicitudin risus. Ut condimentum leo id elementum iaculis. Pellentesque eget facilisis dolor.

Mauris vel mi facilisis, posuere lectus et, pulvinar lorem. Vivamus congue orci ut ullamcorper tempus. Integer a venenatis lectus, in fringilla lorem. In non mi at quam vehicula vulputate. Sed vel lectus ac lacus ornare pharetra vel facilisis urna. Maecenas orci urna, sollicitudin nec velit et, interdum condimentum erat. Praesent nec erat nisi. Nullam ex odio, auctor non vehicula in, tempus non turpis. Proin hendrerit quam a turpis ornare, at tristique turpis lacinia.

Ut sit amet ultricies mi. Donec quis vulputate nisi. Nunc et mi eget ante vehicula pulvinar sit amet nec ex. Vestibulum porttitor pellentesque mauris quis feugiat. Sed efficitur diam eu libero lobortis, quis pretium nisi viverra. Proin blandit est pretium tortor egestas scelerisque. Aenean mauris velit, egestas eu facilisis a, lobortis eget elit. Fusce blandit, metus vel ornare vulputate, metus risus feugiat elit, eu mollis orci ex ac felis. Nullam imperdiet tortor felis, ut sollicitudin turpis ornare id. Nullam finibus neque nec metus ullamcorper posuere. Maecenas pellentesque enim lectus. Curabitur eu dapibus nunc. Interdum et malesuada fames ac ante ipsum primis in faucibus.

Nulla facilisi. Vivamus quis lacus id nulla dapibus molestie eget eu mi. Integer bibendum consectetur dolor, eu facilisis risus venenatis vitae. Aliquam volutpat lacinia nibh, sed viverra enim aliquet ut. Maecenas a nunc pharetra, hendrerit diam non, feugiat sapien. In malesuada at leo sit amet molestie. Sed dui massa, rhoncus ut tincidunt vel, consequat non nibh. Maecenas varius venenatis massa, nec dignissim nisl rhoncus eu. Maecenas euismod dui id orci tempor pharetra. Aenean viverra nulla blandit consectetur pretium. Sed posuere enim sit amet magna laoreet congue eget pretium mauris. Orci varius natoque penatibus et magnis dis parturient montes, nascetur ridiculus mus. Curabitur fermentum varius nisi, et porttitor mauris. Vestibulum ante ipsum primis in faucibus orci luctus et ultrices posuere cubilia Curae; Donec gravida faucibus fermentum.

Vestibulum ante ipsum primis in faucibus orci luctus et ultrices posuere cubilia Curae; Nunc malesuada rutrum placerat. Donec dapibus aliquam dolor eu fermentum. Aenean vestibulum leo mattis, pulvinar felis scelerisque, venenatis sapien. Nulla non maximus lorem, non condimentum augue. Sed vel sagittis risus, sit amet pellentesque libero. Etiam vel ornare risus. Morbi sagittis lacinia justo sed consectetur. Pellentesque sed blandit velit. Donec ligula ligula, luctus sed massa id, dapibus faucibus libero. Suspendisse dui enim, ullamcorper vitae pulvinar in, elementum quis nulla. Cras tempor neque eu nibh aliquet, eu tempus mauris accumsan.

Vestibulum aliquet vel arcu ac efficitur. Praesent placerat augue quis pellentesque porttitor. Suspendisse molestie et sem vel lacinia. Etiam scelerisque, est sagittis ultrices pellentesque, ex ligula fermentum felis, id rhoncus elit nunc porta sem. Nunc velit orci, ultrices non volutpat eget, facilisis ac ex. Pellentesque pulvinar interdum porta. Quisque et lorem id neque consequat cursus sed vitae tellus. Duis posuere, ligula eget sollicitudin ultricies, sapien nisi maximus ligula, a suscipit orci sem et ex. Nam sit amet porta ligula, ac dapibus nisi. Pellentesque habitant morbi tristique senectus et netus et malesuada fames ac turpis egestas.

Duis sapien ligula, porta at leo non, vulputate sagittis leo. Class aptent taciti sociosqu ad litora torquent per conubia nostra, per inceptos himenaeos. Curabitur pretium ante eu lectus commodo fermentum. Duis scelerisque facilisis purus aliquam blandit. In sit amet nunc metus. Nullam condimentum velit at tristique vestibulum. Aenean non feugiat lorem. In tempus ante et velit ultricies vestibulum. Duis consectetur lectus eget arcu congue, vel lobortis urna hendrerit.

Aliquam vitae lacinia ipsum. Ut aliquet ligula ut tincidunt lacinia. Quisque eget sodales quam. Pellentesque aliquet feugiat enim. Aenean a dolor eget elit fermentum feugiat. Nullam ac scelerisque leo. Sed gravida ex odio, eu commodo dui sagittis eget. Mauris ullamcorper pretium lorem, at bibendum neque.

Lorem ipsum dolor sit amet, consectetur adipiscing elit. Morbi molestie luctus libero rhoncus consequat. Vivamus hendrerit finibus erat quis gravida. Pellentesque quis felis sit amet sem maximus commodo nec ac lectus. In hac habitasse platea dictumst. Aenean vel mi convallis, tristique sem ac, varius enim. Integer porta condimentum neque vel aliquam.

Quisque volutpat, lorem sed consectetur interdum, magna lacus pellentesque est, eu eleifend quam elit et risus. Proin auctor efficitur quam, sodales dapibus mauris elementum quis. Maecenas fringilla, nibh ac tempor molestie, neque nulla dapibus urna, eu mattis nunc eros finibus turpis. Nulla sit amet dui ornare, facilisis massa et, malesuada urna. Fusce egestas eros urna, at mollis eros mollis quis. Sed fringilla vestibulum tellus in vehicula. Suspendisse ornare quis enim sed tincidunt. Donec sed sem ultricies, fermentum velit vel, pellentesque nisi. Praesent tincidunt felis eget feugiat aliquam. Maecenas faucibus lectus eros, eu ultricies diam hendrerit ut. Vestibulum vitae pretium urna, vel sollicitudin sapien.

Maecenas sit amet lorem leo. Aenean pellentesque congue lorem, vel rhoncus dolor fermentum id. Donec lacinia ligula at cursus sollicitudin. Proin sed erat velit. Nunc ac tortor at nisi suscipit rutrum. Nam porttitor est a diam tempus iaculis. Aliquam neque est, aliquet ac iaculis sit amet, sodales vitae leo. Phasellus viverra finibus est, quis feugiat metus interdum sit amet. Mauris ac finibus diam, at suscipit ex.

Sed in justo quam. Morbi interdum justo non diam facilisis, ut facilisis odio varius. Vestibulum efficitur leo sed magna porta mattis. Nunc non justo eleifend, tristique felis eget, fermentum neque. Nulla hendrerit nunc nisl, eu fringilla ex euismod varius. Nunc a ultrices risus. Ut non dolor tortor. Aenean eu interdum leo. Duis facilisis iaculis sollicitudin. Aenean rhoncus risus rutrum ipsum volutpat, id tincidunt augue commodo. Donec sed urna efficitur, commodo sem vitae, tincidunt sapien.

Fusce pulvinar urna ut eleifend aliquam. Curabitur nec lacus efficitur, ullamcorper libero vel, maximus nunc. In sapien felis, convallis tincidunt pellentesque quis, molestie vitae libero. Sed nec facilisis urna, nec suscipit velit. In vestibulum justo pellentesque ullamcorper mattis. Integer sed nibh ligula. Integer eleifend laoreet justo, a commodo enim lacinia eget. Sed nec lorem a velit posuere tincidunt. Phasellus id justo tortor. Vestibulum et purus orci. Duis rutrum lectus est, sed luctus nisi suscipit vel.

Proin leo orci, placerat quis eleifend quis, blandit a est. Vestibulum eu libero a mi sagittis porttitor quis in tortor. Maecenas ut ullamcorper neque. Duis in consectetur tortor, quis euismod felis. Cras eu mi in risus pulvinar ultricies. Sed in lorem lacus. Pellentesque at commodo felis. Maecenas vel tincidunt nulla. Quisque aliquam vel nulla ac tempor. Sed pretium eu purus a elementum. Vivamus ut ullamcorper est.

Praesent in commodo ante, eget lacinia ipsum. Sed luctus pharetra sagittis. Proin sed sapien vel lacus semper venenatis eu id quam. Praesent interdum luctus metus sit amet pulvinar. Etiam turpis odio, dapibus vel erat elementum, aliquam auctor nulla. Phasellus pretium est massa, vitae mattis dui commodo at. Quisque dictum sit amet nulla quis condimentum. Quisque varius eros in nibh vulputate placerat. Proin vitae tempor massa. Orci varius natoque penatibus et magnis dis parturient montes, nascetur ridiculus mus. Maecenas pellentesque odio vitae diam luctus mollis. Cras auctor risus vitae nisi ultricies blandit. Nullam condimentum lorem ut erat suscipit viverra. Nunc tincidunt nisl amet.`

var sentenceMu sync.Mutex
var sentenceBuf = bytes.NewBuffer(nil)
var sentenceIdx int

func GetSentence() (sentence string, index int) {
	sentenceMu.Lock()
	defer sentenceMu.Unlock()
	for {
		var err error
		sentence, err = sentenceBuf.ReadString('.')
		if err != nil {
			if err == io.EOF {
				sentenceBuf = bytes.NewBufferString(Lipsum)
				sentenceIdx = 0
				continue
			}
			panic(err)
		}
		index = sentenceIdx
		sentenceIdx++
		break
	}
	return strings.TrimSpace(sentence), index
}
