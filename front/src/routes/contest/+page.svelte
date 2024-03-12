<script lang="ts">
  import SveltyPicker from 'svelty-picker';

  let valueDefault: string | string[] | undefined =  new Date().toLocaleDateString();
  let valueDifferent: string | string[] | undefined =  new Date().toLocaleDateString();
  let valueTimeOnly: string | string[] | undefined =  new Date().toLocaleDateString();
</script>

<div class="flex items-start gap-8">
  <div>
    <h4>
      Format === displayFormat<br>Internal value: <code>{valueDefault}</code>
    </h4>
    <SveltyPicker bind:value={valueDefault} />
  </div>
  <div>
    <h4>
      Format !== displayFormat<br>Internal value: <code>{valueDifferent}</code>
    </h4>
    <SveltyPicker
      bind:value={valueDifferent}
      displayFormat="j. n. Y"
      displayFormatType="php"
    />
  </div>
  <div>
    <h4>
      Time only format:<br>Internal value: <code>{valueTimeOnly}</code>
    </h4>
    <SveltyPicker bind:value={valueTimeOnly} format="hh:ii" displayFormat="HH:ii P" />
  </div>
</div>
