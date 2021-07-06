const axios = require('axios');

export const SettingsService = {
	fetchSettings: () => {
		return axios.get('http://localhost:8085/api/libify/settings-v1')
			.then(res => res.data);
	},

	save: settings => {
		return axios.put(`http://localhost:8085/api/libify/settings-v1`, settings)
	},
};
