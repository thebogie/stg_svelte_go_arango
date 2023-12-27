import { redirect } from '@sveltejs/kit';
import { loginPlayer } from '$lib/services/player.service';
import type { IPlayer } from '$lib/interfaces/player.d';

export const actions = {
	login: async ({ cookies, request }) => {
		console.log('LOGIN ACTIONS');

		let token = 'EMPTY';
		let player: IPlayer = {};

		const data = await request.formData();
		let emaildata = data.get('email');
		let passworddata = data.get('password');

		if (emaildata == null) {
			return;
		}
		if (passworddata == null) {
			return;
		}

		try {
			let player = await loginPlayer(emaildata.toString(), passworddata.toString());
			console.log('logged_in_player' + JSON.stringify(player));

			if (!player.accessToken) {
				return;
			}
			token = player.accessToken;

			cookies.set('player', JSON.stringify(player), {
				httpOnly: true,
				secure: true,
				path: '/'
			});
		} catch (err) {
			console.error(err); // Add your error handling here
		}

		throw redirect(303, '/profile');
	}
};
