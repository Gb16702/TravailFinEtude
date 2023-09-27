<script lang="ts">
    import SelectCountry from "../../../components/SelectCountry.svelte";
    import type { PageData } from "./$types";
    import Search from "../../../components/icons/Search.svelte";
    import Close from "../../../components/icons/Close.svelte";
    import Checkbox from "../../../components/Checkbox.svelte";

    type Checkbox = {
        id: string,
        checked: boolean
    }

    export let data:PageData;
    export let open = true;

    let search: string = "";
    let countryNames: any = data.countries.map((item: Record<string, string>) => {
        return {
            name: item.name,
            code: item.code
        }
    });

    const handleChange: (event: any) => Record<string, string> = event => {
        if(event && event.target) {
            const { value } = event.target as HTMLInputElement;
            search = value
            return countryNames = data.countries
            .filter((obj: Record<string, string>) => obj.name.toLowerCase().includes(value.toLowerCase()) as boolean)
            .map((item: Record<string, string>) => {
                return {
                    name: item.name,
                    code: item.code
                }
            });
        }
    }

    const resetSearch = (): void => {
        search = "";
        countryNames = data.countries.map((item: Record<string, string>) => {
            return {
                name: item.name,
                code: item.code
            };
        })
    }

    let checkbox: Checkbox = Object.assign({id: "", checked: false})
    const handleCheck: (id:string) => void = (id) => {
        checkbox = {id, checked: !checkbox.checked}
    };
</script>

<div class="w-full h-screen bg-[#ccc]">
    <SelectCountry countryData countryProps={countryNames} />
    {#if open}
        <div class="fixed top-1/2 left-1/2 -translate-x-1/2 -translate-y-1/2 w-[90%] bg-white rounded-[8px] overflow-hidden">
            <div class="w-full h-[60px] bg-white px-5 flex flex-row justify-between items-center">
                    <h3 class="font-semibold">
                        Sélectionner un pays
                    </h3>
                <div class="bg-zinc-100 p-[2px] rounded-md font-semibold text-[15px] cursor-pointer" on:click={() => open = false}>
                    <Close className="w-[16px] h-[16px] stroke-zinc-500" />
                </div>
            </div>
            <div class="w-full h-[60px]  flex items-center justify-center px-5">
                <div class="w-full flex flex-row gap-x-[2px] bg-white items-center justify-between h-[33px] rounded-md overflow-hidden">
                    <div class="flex flex-row items-center gap-x-2">
                    <div class="w-[10%] max-w-[27px] h-[100%] flex items-center justify-start">
                        <Search className="w-[18px] h-[18px] stroke-zinc-400" />
                    </div>
                    <input type=""
                    placeholder="Rechercher..."
                    class="w-[80%] h-full rounded-[5px] text-[15px] outline-none"
                    on:input={handleChange}
                    bind:value={search}
                    >
                </div>
                {#if search.trim().length !== 0 }
                    <div class="w-[10%] max-w-[27px] h-[100%] flex items-center justify-end" on:click={resetSearch}>
                            <Close className="w-[12px] h-[12px] stroke-zinc-500" />
                    </div>
                {/if}
                </div>
            </div>
            <div class="w-full bg-white overflow-auto h-[240px] pt-1 pb-5">
                {#if countryNames.length !== 0}
                    {#each countryNames as cn }
                        <div class={`w-full h-[45px] flex items-center ${checkbox.id == cn && checkbox.checked && "bg-zinc-100"}`} on:click={() => handleCheck(cn)}>
                            <div class="flex justify-between flex-row px-5 items-center w-full">
                                <div class="flex flex-row gap-x-2 items-center">
                                    <Checkbox theme="alt" />
                                    <h2>
                                        {cn.name}
                                    </h2>
                                </div>
                                <h2>
                                    {cn.code}
                                </h2>
                            </div>
                        </div>
                        {/each}
                    {:else}
                    <div class="w-full h-full flex items-center justify-center text-center px-2 font-medium text-[15px]">
                        Aucun résultat pour la recherche : {search}
                    </div>
                    {/if}
            </div>
        </div>
    {/if}
    <div class="w-[300px] flex flex-row h-[50px] mt-3">
        <div class="bg-red-400 w-[20%]" on:click={
            () => {
                open = !open
            }
        }>
        </div>
        <input
        on:input={handleChange}
        />
    </div>
</div>

