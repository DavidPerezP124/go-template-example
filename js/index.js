
// Playing logic
const setPlaying = (e, siblings) => {
	// Turn of other players
	for (let idx = 0; idx < siblings.length; idx++) {
		const sibling = siblings[idx]
		if (e.currentTarget === sibling) continue;
		sibling.classList.remove("is-playing")
		sibling.querySelector(".audio-player").pause()
	}
	// Toggle the current player
	const target = e.currentTarget
	target.classList.toggle("is-playing")
	if (target.classList.contains('is-playing')) {
		e.currentTarget.querySelector(".audio-player").play()
	} else {
		e.currentTarget.querySelector(".audio-player").pause()
	}
	addListeners(e.currentTarget.querySelector(".audio-player"),target)
}

const albums = document.getElementsByClassName("music-player-container")

for (let idx = 0; idx < albums.length; idx++) {
	albums[idx].onclick = (e) => setPlaying(e, albums)
}

const addListeners = (video, parent) => {
	video.onended = (element) => {
		parent.classList.remove("is-playing")
	}
}