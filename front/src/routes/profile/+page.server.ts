import { redirect } from '@sveltejs/kit';

export function load(event) {
	console.log("Profile +page.server.ts load");
	if ( event.locals.player == null ) {
		throw redirect(303, '/login');
	}
}