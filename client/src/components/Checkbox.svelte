<script lang="ts">
    export let theme: "default" | "alt" = "default";
    export let checkedClasses: string = "";

    let checked: boolean = false;
    let defaultClasses: string = "absolute -translate-x-1/2 -translate-y-1/2 pointer-events-none top-1/2 left-1/2 w-full h-full flex items-center justify-center overflow-hidden rounded-[4px]"

    $: {
        if(theme == "default") {
            checkedClasses="bg-accent-violet-d"
        } else if(theme == "alt") {
            checkedClasses="bg-black"
        }
    }

</script>
<div class="relative w-fit h-fit overflow-hidden border rounded-[4px]">
    <div class="appearance-none rounded-[4px] bg-white overflow-hidden w-4 h-4" on:click={() => {
        checked = !checked;
    }}></div>
    {#if checked}
    <div class={`${defaultClasses} ${checkedClasses}`}>
        <svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke-width="3" stroke="white" class="w-[80%] h-[80%]">
        <path stroke-linecap="round" stroke-linejoin="round" d="M4.5 12.75l6 6 9-13.5" />
    </svg>
</div>
{/if}
</div>

<style>
    .checkbox{
        appearance: none;
        border-radius: 4px;
        background-color: white;
        overflow: hidden;
    }
</style>