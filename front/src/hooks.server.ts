import { checkPlayer } from '$lib/services/player.service';

export async function handle({ event, resolve }) {
	const token = event.cookies.get('token');

	console.log('EVENT: ' + JSON.stringify(event));
	if (token == undefined) {
		return await resolve(event);
	}

	try {
		let playerCookie = event.cookies.get('player');
		if (!playerCookie) {
			return await resolve(event);
		}

		const player = JSON.parse(playerCookie);

		event.locals.player = player;
		//event.locals.player = checkPlayer()
	} catch (error) {
		console.log(error);
	}

	const response = await resolve(event);
	return response;
}
