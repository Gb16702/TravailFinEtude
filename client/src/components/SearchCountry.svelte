<script lang="ts">
    import { countryStore } from "../stores/countryStore";
    import { isOpen } from "../stores/openStore";
    import Checkbox from "./Checkbox.svelte";
    import Close from "./icons/Close.svelte";
    import Search from "./icons/Search.svelte";

    type Checkbox = {
        id: number | null,
        checked: boolean,
        code: string | null
    };

    export let data: any;

    let search: string = "";
    let countryNames: any = data.countries
        .map((item: Record<string, string>) => Object.assign({
                id: item.id,
                name: item.name,
                code: item.country_code,
                phone_number: item.phone_number,
                logo: item.logo
            })
    );

    const handleClose = (): boolean => {
        return $isOpen = false
    };

    const handleChange: (event: any) => Record<string, string> = event => {
        if(event && event.target) {
            const { value } = event.target as HTMLInputElement;
            search = value
            return countryNames = data.countries
            .filter((obj: Record<string, string>) => obj.name.toLowerCase().includes(value.toLowerCase()) as boolean)
            .map((item: Record<string, string>) => Object.assign({
                    id: item.id,
                    name: item.name,
                    code: item.country_code,
                    phone_number: item.phone_number,
                    logo: item.logo
                })
            );
        };
    };

    const resetSearch = (): void => {
        search = "";
        countryNames = data.countries.map((item: Record<string, string>) => {
            return {
                id: item.id,
                name: item.name,
                code: item.country_code,
                phone_number: item.phone_number,
                logo: item.logo
            };
        });
    };

    let checkbox: Checkbox = Object.assign({id: null, checked: false, code: null});
    const handleCheck = (id: number | null, phone_number: string | null): Checkbox => {
        countryStore.set(phone_number);
        return checkbox.id === id ? checkbox = {id: null, checked: false, code: null } : checkbox = {id, checked: true, code: phone_number}
    };
</script>

    {#if $isOpen}
        <div class="fixed z-1 bg-black/[.2] backdrop-blur-sm w-full h-full top-0 left-0"></div>
        <div class="fixed z-10 top-1/2 left-1/2 -translate-x-1/2 -translate-y-1/2 w-[90%] bg-white rounded-[8px] overflow-hidden">
            <div class="w-full h-[60px] bg-white px-5 flex flex-row justify-between items-center">
                    <h3 class="font-semibold">
                        Sélectionner un pays
                    </h3>
                <div class="bg-zinc-100 p-[2px] rounded-md font-semibold text-[15px] cursor-pointer" on:click={handleClose} >
                    <Close className="w-[16px] h-[16px] stroke-zinc-500" />
                </div>
            </div>
            <div class="w-full h-[53px]  flex items-center justify-center px-5">
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
            <div class="w-full bg-white overflow-auto h-[220px] pt-1 pb-5">
                {#if countryNames.length !== 0}
                    {#each countryNames as cn }
                        <div class={`w-full h-[45px] flex items-center ${checkbox.id == cn && checkbox.checked && "bg-zinc-100"}`} on:click={e => handleCheck(cn.id, cn.phone_number)}>
                            <div class="flex justify-between flex-row px-5 items-center w-full">
                                <div class="flex flex-row gap-x-2 items-center">
                                    <Checkbox checked={checkbox.id === cn.id}  theme="alt" />
                                    <h2>
                                        {cn.name}
                                    </h2>
                                </div>
                                <div class="h-5 w-5 text-[20px] flex items-center justify-center">{cn.logo}</div>
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