import { writable, get } from 'svelte/store';
import type { IPlayer } from '$lib/interfaces/player.interface';
import { getSharedStore } from '$lib/stores/use-shared-store';

export const playerStore = () => {
	const { set, update, subscribe } = writable<IPlayer>();
	return {
		set,
		update,
		subscribe
	};
};

export const getPlayerStore = () => getSharedStore('logged_in', playerStore);
