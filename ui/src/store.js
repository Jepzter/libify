import { createStore } from 'vuex'

export default createStore({
	state: {
		currentPlaylistId: undefined,
		currentSong: {},
		landing: true,
	},
	mutations: {
		setCurrentPlaylistId: (state, playlistId) => {
			state.currentPlaylistId = playlistId;
		},
		setCurrentSong: (state, song) => {
			state.currentSong = song;
		},
		setLanding: (state, landing) => {
			state.landing = landing;
		}
	},
	actions: {
		setCurrentPlaylistId: ({ commit }, playlistId) => {
			commit('setCurrentPlaylistId', playlistId);
		},
		setCurrentSong: ({ commit }, song) => {
			commit('setCurrentSong', song);
		},
		setLanding: ({ commit }, landing) => {
			commit('setLanding', landing);
		}
	}
})
