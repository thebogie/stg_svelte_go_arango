import {useWritable } from '$lib/stores/use-shared-store';
import {writable} from 'svelte/store';
import type {IProfile} from '$lib/interfaces/contest';

export const profileStore = () => {
  const { set, update, subscribe } = writable<IProfile>()
  return {
    set,
    update,
    subscribe,
    setProfile: (profile: IProfile) =>set(profile)
  }
};

export const getProfile = () =>
  useWritable('profile', profileStore());
