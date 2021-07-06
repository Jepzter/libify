export const SongService = {
	fetchSongsByPlaylistId: (playlistId) => {
		const songs = [
			{
				name: "song 1",
				artist: "artist 1",
				duration: "4:00",
				location: "test.mp3",
				id: 1,
			},
			{
				name: "song 2",
				artist: "artist 2",
				duration: "3:14",
				location: "test1.mp3",
				id: 2,
			},
			{
				name: "song 3",
				artist: "artist 3",
				duration: "2:13",
				location: "test.mp3",
				id: 3,
			},
		];
		if (playlistId == 2) {
			return [
				...songs,
				{
					name: "song 4",
					id: 4,
					artist: "artist 4",
					duration: "3:14",
					location: "test.mp3",
				},
			];
		}
		return songs;
	}
};
