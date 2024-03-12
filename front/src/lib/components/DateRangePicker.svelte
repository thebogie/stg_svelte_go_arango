<script lang="ts">
  import SveltyPicker from 'svelty-picker';

  type Option = {
    value: Date;
    label: string;
  };

  interface Props {
    startDate: Option | null;
    endDate: Option | null;
    onRangeChange: (start: Date, end: Date) => void;
  }

  const now = new Date();

  let start: Date = now;
  let end: Date = now;

  function handleStartDateChange(newStart: Date) {
    if (newStart > end) {
      end = newStart;
    }
    start = newStart;
    props.onRangeChange(start, end);
  }

  function handleEndDateChange(newEnd: Date) {
    end = newEnd;
    props.onRangeChange(start, end);
  }
</script>

<svelte:head>
  <link href="https://unpkg.com/svelty-picker@^1.13.1/dist/svelty-picker.min.css" rel="stylesheet" />
</svelte:head>

<div>
  <label for="startDate">Start Date & Time:</label>
  <SveltyPicker
    format="yyyy-MM-dd HH:mm"
    name="startDate"
    onInputChange={handleStartDateChange}
    value={start}
  />

  <label for="endDate">End Date & Time:</label>
  <SveltyPicker
    format="yyyy-MM-dd HH:mm"
    minDate={start}
    name="endDate"
    onInputChange={handleEndDateChange}
    value={end}
  />
</div>
