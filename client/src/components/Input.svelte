<script lang="ts">
    import { isOpen } from "../stores/openStore";
    import Arrow from "./icons/Arrow.svelte";
    import { countryStore } from "../stores/countryStore";

    type Input = {
        input_type: "text" | "password" | "email" | "number" | "tel" | "hidden" | null;
        name: string;
        placeholder: string;
        label?: string | null;
        value: string | boolean;
        additionalClasses: string;
        tag?: string;
        passwordClasses?: string;
        IsPasswordInput: boolean;
    }

    const handleOpen = () => isOpen.set(!$isOpen);

    console.log($countryStore.code);

    export let type: Input["input_type"] = "text";
    export let name: Input["name"] = "";
    export let placeholder: Input["placeholder"] = "";
    export let label: Input["label"] = null;
    export let value: Input["value"] = "fsqddqsd";
    export let additionalClasses: Input["additionalClasses"] = "";
    export let passwordClasses: Input["passwordClasses"] = "";

    let inputBaseClass="w-full h-full px-2 overflow-hidden placeholder-zinc-500/[.8] outline-none text-[14.5px] bg-transparent";
    let classes = "rounded-[6px] overflow-hidden h-[40px] transition-all bg-white border border-zinc-200/[.7] duration-200 focus-within:outline-zinc-300 px-2 ";
</script>

<div class={`${label && "flex flex-col gap-y-1"} ${additionalClasses}`}>
    {#if label}
        <label for={name} class="text-[13px] tracking-wide text-zinc-600 mt-2.5">{label}</label>
    {/if}
    <div class={`${passwordClasses} ${classes}`}>
        {#if type === "text" }
            <input
                name={name}
                bind:value={value}
                type="text"
                class={inputBaseClass}
            />
            <slot />
        {:else if type==="email"}
            <input
                name={name}
                class={inputBaseClass}
                bind:value={value}
                type="email"
            />
        {:else if type==="password"}
            <input
                name={name}
                class={inputBaseClass}
                bind:value={value}
                placeholder={placeholder}
                type="password"
            />
            <slot />
        {:else if type==="tel"}
            <div class="flex flex-row h-full">
                <div class="w-[70px] h-full border-r border-zinc-200/[.7] flex flex-row items-center" on:click={handleOpen}>
                    <div class="w-1/2 flex items-center justify-center">
                        <div class="w-5 h-5 flex items-center justify-center ">
                            {#if $countryStore.logo !== null}
                            {$countryStore.logo}
                            {/if}
                        </div>
                    </div>
                    <div class="w-1/2 flex items-center justify-center">
                        <div class="w-5 h-5 flex items-center justify-center">
                            <Arrow classes="w-[14px] h-[14px] relative top-[2px] stroke-zinc-600" stroke_width="2" />
                        </div>
                    </div>
                </div>
                <div class="w-[55px] h-full flex items-center justify-end text-[14.5px]">
                    {#if $countryStore.code !== null}
                            {$countryStore.code}
                    {/if}
                </div>
                <input
                    name={name}
                    bind:value={value}
                    type="tel"
                    class={inputBaseClass}
                />
            </div>
        {/if}
    </div>
</div>
