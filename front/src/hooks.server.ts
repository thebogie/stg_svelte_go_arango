
import { redirect } from '@sveltejs/kit';

export async function handle({ event, resolve }) {
	console.log('HANDLE: ' + JSON.stringify(event));
	let playerCookie = event.cookies.get('player');

	if (!playerCookie) {
		if (event.route.id !== '/login') {
			throw redirect(303, '/login');
		}
	} else {
		event.locals.player = JSON.parse(playerCookie);
	}

	const response = await resolve(event);
	return response;
}
