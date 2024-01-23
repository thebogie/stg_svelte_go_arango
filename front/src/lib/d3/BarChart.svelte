<script lang="ts">
  import {page} from '$app/stores';
  import {scaleBand, scaleLinear} from 'd3';
  import type {IContest} from '$lib/interfaces/contest';
  import {countResults} from '$lib/utils/common';


  //console.log('BARCHART: ' + JSON.stringify($page.data));
  //let data = $page.data.countries;
  let total_results: IContest[] = $page.data.total_results;

  let win_loss = countResults(total_results);
  console.log('winlose: ' +win_loss);

  let width = 800;
  let height = 200;

  const margin = {top: 20, right: 20, bottom: 20, left: 180};
  const innerHeight = height - margin.top - margin.bottom;
  const innerWidth = width - margin.left - margin.right;

  $: xDomain = win_loss.map((d) => d.status);
  $: yDomain = win_loss.map((d) => +d.value);

  $: yScale = scaleBand().domain(xDomain).range([0, innerHeight]).padding(0.1);
  $: xScale = scaleLinear()
    .domain([0, Math.max.apply(null, yDomain)])
    .range([0, innerWidth]);
</script>

<svg {height} {width}>
  <g transform={`translate(${margin.left},${margin.top})`}>
    {#each xScale.ticks() as tickValue}
      <g transform={`translate(${xScale(tickValue)},0)`}>
        <line y2={innerHeight} stroke="green" />
        <text text-anchor="middle" dy=".7em" y={innerHeight + 3}>
          {tickValue}
        </text>
      </g>
    {/each}
    {#each win_loss as d}
      <text
        text-anchor="end"
        x="-3"
        dy=".3em"
        y={yScale(d.status) + yScale.bandwidth() / 2}
      >
        {d.status} : ({d.value})
      </text>
      <rect
        x="0"
        y={yScale(d.status)}
        width={xScale(d.value)}
        height={yScale.bandwidth()}
      />

    {/each}
  </g>
</svg>
