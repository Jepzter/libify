const axios = require('axios');

export const SongService = {
	findSongs: (search) => {
		return axios.get(`http://localhost:8085/api/libify/song-v1?Title=${search}`)
			.then(res => res.data);
	}
};
