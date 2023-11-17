/** @type {import('./$types').Actions} */

import type { PageServerLoad, Actions } from './$types';
import { redirect, fail } from '@sveltejs/kit';
import { loginPlayer } from '$lib/services/player.service';
import type { IPlayer } from '$lib/interfaces/player.interface';

export const load: PageServerLoad = (event) => {
	const user = event.locals.player;

	if (user) {
		throw redirect(302, '/profile');
	}
};

export const actions: Actions = {
	default: async (event) => {
		const formData = Object.fromEntries(await event.request.formData());

		if (!formData.email || !formData.password) {
			return fail(400, {
				error: 'Missing email or password'
			});
		}

		const { email, password } = formData as {
			email: string;
			password: string;
		};

		let loggedinplayer: IPlayer = {};

		try {
			loggedinplayer = await loginPlayer(email, password);

			// Set the cookie
			event.cookies.set('Authorization', `${loggedinplayer.accessToken}`, {
				httpOnly: true,
				path: '/',
				secure: true,
				sameSite: 'strict',
				maxAge: 60 * 60 * 24 // 1 day
			});

			event.cookies.set('loggedinplayer', JSON.stringify(loggedinplayer), {
				httpOnly: true,
				path: '/',
				secure: true,
				sameSite: 'strict',
				maxAge: 60 * 60 * 24 // 1 day
			});

			event.locals.player = loggedinplayer;
		} catch (err: any) {
			console.log('login error: ' + err.message);
			return fail(401, {
				error: err.message
			});
		}

		throw redirect(302, '/profile');
	}
};
