<script lang="ts">
    export let countryData;
    export let open:boolean=false;

    export let countryProps
</script>

<button on:click={() => {
    open= !open
}} class="w-[100px] h-[100px] bg-red-400 rounded-10px">
    {#if open}
    Fermer
    {:else}
    Ouvrir
    {/if}
</button>
{#if open}
<div>
    {countryProps}
</div>
{/if}

<style>


</style>