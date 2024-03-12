import { redirect } from '@sveltejs/kit';
import type { Actions, PageServerLoad } from './$types';

export const load: PageServerLoad = async () => {
	throw redirect(302, '/');
};

export const actions: Actions = {
	default({ cookies }) {
		console.log('logout');
		cookies.set('player', '', {
			path: '/',
			expires: new Date(0)
		});

		throw redirect(302, '/');
	}
};
