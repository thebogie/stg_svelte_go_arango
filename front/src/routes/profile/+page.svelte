<script lang="ts">
    import BarChart from "$lib/d3/BarChart.svelte"
    import Cookies from "js-cookie";
    import {getCurrentPlayer} from '$lib/stores/playerStore';
    import {getContext} from "svelte";
    import type {IPlayer} from "$lib/interfaces/player.interface";
    //let count = countResults(data);


    let data = [
        {country: "China", population: 1439324},
        {country: "India", population: 1380004},
        {country: "United States of America", population: 331003},
        {country: "Indonesia", population: 273524},
        {country: "Pakistan", population: 220892},
        {country: "Brazil", population: 212559},
        {country: "Nigeria", population: 206140},
        {country: "Bangladesh", population: 164689},
        {country: "Russian Federation", population: 145934},
        {country: "Mexico", population: 128933},
    ];
</script>


<svelte:head>
    <title>Profile</title>
</svelte:head>


<div class="card p-4">Games Played

    <BarChart {data}/>

</div>
<h1>Profile</h1>
<p></p>


