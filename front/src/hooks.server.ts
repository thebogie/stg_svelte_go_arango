import { redirect } from '@sveltejs/kit';

export async function handle({ event, resolve }) {
	console.log('HANDLE: ' + event.route.id);
	let playerCookie = event.cookies.get('player');

	if (!playerCookie) {
		if (event.route.id !== '/login' && event.route.id !== '/') {
			throw redirect(303, '/login');
		}
	} else {
		event.locals.player = JSON.parse(playerCookie);
	}
	/*
    if (event.route.id !== '/logout' ) {
      event.cookies.set('player', '', {
        path: '/',
        expires: new Date(0)
      });
      //throw redirect(303, '/');
    }
  
  
   */
	const response = await resolve(event);
	return response;
}
