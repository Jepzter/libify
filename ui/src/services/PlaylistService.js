const axios = require('axios');

export const PlaylistService = {
	fetchPlaylists: () => {
		return axios.get('http://localhost:8085/api/libify/playlist-v1')
			.then(res => res.data);
	},

	fetchPlaylistById: playlistId => {
		return axios.get(`http://localhost:8085/api/libify/playlist-v1/${playlistId}`)
			.then(res => res.data);
	},
};
