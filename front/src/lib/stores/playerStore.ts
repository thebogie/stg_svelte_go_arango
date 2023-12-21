import { writable, get } from 'svelte/store';
import type { IPlayer } from '$lib/interfaces/player.interface';
import { loginPlayer } from '$lib/services/player.service';
import Cookies from 'js-cookie';
import { goto } from '$app/navigation';

export const logged_in_player = writable<IPlayer | null>(null);

export async function loginCurrentPlayer(email: string, password: string): Promise<void> {
	let player;
	console.log('player Store');
	try {
		player = (await loginPlayer(email, password)) as IPlayer;


		if (!player) {
			throw new Error('Login failed:');
		}
		logged_in_player.set(player); // Set the user data in the store
	} catch (error) {
		console.error('Login failed: ', error);
		// Add more error handling here if needed
	}
}

// function to get current user
export function getCurrentPlayer() {
	console.log('getcurrentplayer' + get(logged_in_player));
	let results;
	//if (!Cookies.get('token') || !get(logged_in_player)) {
	//	goto('/login');
	//}

	return results;
}
