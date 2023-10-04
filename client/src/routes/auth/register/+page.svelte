<script lang="ts">
    import { enhance } from "$app/forms";
    import type { SubmitFunction } from "@sveltejs/kit";
    import AuthSectionToggler from "../../../components/AuthSectionToggler.svelte";
    import Button from "../../../components/Button.svelte";
    import Checkbox from "../../../components/Checkbox.svelte";
    import Input from "../../../components/Input.svelte";
    import SearchCountry from "../../../components/SearchCountry.svelte";
    import Hide from "../../../components/icons/Hide.svelte";
    import Show from "../../../components/icons/Show.svelte";
    import type { ActionData, PageData, Snapshot } from "./$types";
	import { countryStore } from './../../../stores/countryStore';
    import GenerateSecurePassword from "../../../utils/generateSecurePassword";
    import Sparkles from "../../../components/icons/Sparkles.svelte";
    import Stepper from "../../../components/Stepper.svelte";

    export const snapshot: Snapshot = {
        capture: () => formDataObject,
        restore: data => formDataObject = data
    }

    const togglePasswordVisibility: () => "text" | "password" = () => {
        return formDataObject.password.type = formDataObject.password_confirm.type = formDataObject.password.type === "password" ? "text" : "password"
    };

    let resultValue: any = null;
    let step: number = 1;

    const handleSubmit:SubmitFunction = () => {
        return async ({result, update}) => {
            resultValue = result;
            step = resultValue?.data?.step
            await update();
        }
    }
    export let form:ActionData;

    let formDataObject: FormDataObject = {
        firstName: {
            type: "text",
            name:"firstname",
            placeholder: "Votre prénom",
            label: "Prénom",
            value: "",
        },
        lastName: {
            type: "text",
            name:"lastname",
            placeholder: "Votre nom",
            label: "Nom",
            value: "",
        },
        email: {
            type: "email",
            name:"email",
            placeholder: "Votre adresse mail",
            label: "Adresse mail",
            value: "",
        },
        phone_number: {
            type:"tel",
            name:"phone_number",
            label: "Numéro de téléphone",
            value: "",
        },
        password: {
            type: "password",
            name:"password",
            label: "Mot de passe",
            value: "",
        },
        password_confirm: {
            type: "password",
            name:"passwordconfirm",
            label: "Confirmer le mot de passe",
            value: "",
        },
        terms: {
            name: "terms",
            type: "hidden",
            value: false
        }
    }

    export let data: PageData

    const generateSecurePassword: () => void = () => {
        formDataObject.password.value = GenerateSecurePassword(12);
        formDataObject.password_confirm.value = formDataObject.password.value;
    }

    const toggleTerms = () => {
        formDataObject.terms.value = !formDataObject.terms.value;
    }

</script>

<div class="w-full flex items-start justify-center flex-col gap-y-[2px] h-[30%]">
    <div class="flex items-center justify-center flex-col w-full h-[55%]">
        <h1 class="text-[19px] text-black">Bienvenue</h1>
        <h1 class="font-bold text-[27px] text-black">Créer un compte</h1>
    </div>
    <Stepper step={step} />
</div>
    <div class="w-full flex flex-col items-center justify-center min-h-[480px]">
        <div class="w-[99%] md:w-[75%] lg:w-[650px] px-2">
        {#if step === 3}
                <form action="?/register" class="rounded-[5px] overflow-hidden" method="POST" use:enhance={handleSubmit}>
                    <input type="hidden" name="step" value=3>
                        <div class="flex flex-col ">
                        <section class=" pb-7 pt-1">
                            <div class=" pl-4 pr-1 py-[20px] rounded-md border border-zinc-300/[.7] bg-white">
                                <div class="pb-4">
                                    <h2 class="pb-1 text-[19px] font-semibold tracking-[-0.010em]">Résumé d'inscription</h2>
                                    <h4 class="text-[13px] text-zinc-400">Les champs peuvent être <span class="font-semibold text-black">modifiés</span> avant de vous inscrire</h4>
                                </div>
                                <div class="flex flex-row w-full h-[35px] mt-4">
                                    <div class="w-[30%] h-full text-sm flex items-center">
                                        <label for="firstName" class="font-medium">
                                            Prénom :
                                        </label>
                                    </div>
                                    <input
                                        class="px-2 overflow-visible rounded-[4px] w-[65%] bg-transparent text-sm placeholder-zinc-500/[.8] outline outline-1 outline-zinc-300/[.7]"
                                        name={formDataObject.firstName.name}
                                        type={formDataObject.firstName.type}
                                        value={formDataObject.firstName.value}
                                    />
                                </div>
                                <div class="flex flex-row w-full mt-4 h-[35px]">
                                    <div class="w-[30%] h-full text-sm flex items-center">
                                        <label for="lastname" class="font-medium">
                                            Nom :
                                        </label>
                                    </div>
                                    <input
                                    class="px-2 rounded-[4px] w-[65%] bg-transparent text-sm placeholder-zinc-500/[.8] outline outline-1 outline-zinc-300/[.7]"
                                    name={formDataObject.lastName.name}
                                    type={formDataObject.lastName.type}
                                    value={formDataObject.lastName.value}
                                    />
                                </div>
                                <div class="flex flex-row w-full mt-4 h-[35px]">
                                    <div class="w-[30%] h-full text-sm flex items-center">
                                        <label for="email" class="font-medium">
                                            Adresse mail :
                                        </label>
                                    </div>
                                    <input
                                    class="px-2 rounded-[4px] w-[65%] bg-transparent text-sm placeholder-zinc-500/[.8] outline outline-1 outline-zinc-300/[.7]"
                                    name={formDataObject.email.name}
                                    type={formDataObject.email.type}
                                    value={formDataObject.email.value}
                                    />
                                </div>
                                <div class="flex flex-row w-full mt-4 h-[35px]">
                                    <div class="w-[30%] h-full text-sm flex items-center">
                                        <label for="password" class="font-medium">
                                            Mot de passe :
                                        </label>
                                    </div>
                                        <div class="px-2 rounded-[4px] w-[65%] bg-transparent  text-sm placeholder-zinc-500/[.8] outline-1 outline-zinc-300/[.7] outline flex flex-row items-center">
                                            <input
                                            class="w-[90%] outline-none"
                                            name={formDataObject.password.name}
                                            type={formDataObject.password.type}
                                            value={formDataObject.password.value}
                                            />
                                            <button class="px-1" on:click|preventDefault={togglePasswordVisibility}>
                                                {#if formDataObject.password.type==="password"}
                                                    <Show className="w-5 h-5 stroke-zinc-500    " />
                                                {:else if formDataObject.password.type==="text"}
                                                    <Hide className="w-5 h-5" />
                                                {/if}
                                            </button>
                                        </div>
                                    </div>
                                <div class="flex flex-row w-full mt-4 h-[35px]">
                                    <div class="w-[30%] h-full text-sm flex items-center">
                                        <label for="tel" class="font-medium">
                                            Téléphone :
                                        </label>
                                    </div>
                                    <input
                                        class="px-2 rounded-[4px] w-[65%] bg-transparent text-sm placeholder-zinc-500/[.8] outline outline-1 outline-zinc-300/[.7]"
                                        name={formDataObject.phone_number.name}
                                        type={formDataObject.phone_number.type}
                                        value={$countryStore.code && $countryStore.code + formDataObject.phone_number.value}
                                    />
                                </div>
                            </div>
                        </section>
                        <section class="flex row w-full flex-1 min-h-[70px] rounded-[5px] border border-zinc-300/[.7] bg-white">
                            <div class=" w-[10%] flex items-center justify-center">
                                <input type={formDataObject.terms.type} name={formDataObject.terms.name} value={formDataObject.terms.value}>
                                <Checkbox theme={"default"} onCheck={toggleTerms} checked={Boolean(formDataObject.terms.value)} />
                            </div>
                            <div class="w-[90%] flex items-center  pr-1">
                                <small class="text-[14px] text-zinc-600">
                                    En cochant cette case, vous acceptez les
                                    <span class="text-black font-medium">
                                        Conditions Générales d'Utilisation
                                    </span>
                                </small>
                            </div>
                        </section>
                    </div>
                    <div class="w-full md:flex md:justify-between md:items-center md:bg-white mt-4 md:px-2 md:rounded-md md:border md:border-zinc-300/[.7] md:py-3">
                        <div class="max-md:hidden">
                            <AuthSectionToggler />
                        </div>
                        <Button additionalClasses="w-full ">
                            Continuer
                        </Button>
                    </div>
                </form>
        {:else if step === 2}
                <form action="?/register" method="POST" use:enhance={handleSubmit}>
                    <input type="hidden" name="step" value=2>
                    <div class="flex flex-col sm:gap-x-1.5 w-full">
                            <Input type={formDataObject.password.type} name={formDataObject.password.name} label={formDataObject.password.label} bind:value={formDataObject.password.value} additionalClasses="w-full"
                                passwordClasses="flex flex-row justify-between items-center">
                                <SearchCountry data={data} />
                                <div class="flex items-center justify-between">
                                    <button class="px-[6px] flex items-center justify-center" on:click|preventDefault={generateSecurePassword}>
                                        <Sparkles  />
                                    </button>
                                    <button class="px-1" on:click|preventDefault={togglePasswordVisibility}>
                                        {#if formDataObject.password.type==="password"}
                                            <Show className="w-5 h-5 stroke-zinc-500    " />
                                        {:else if formDataObject.password.type==="text"}
                                            <Hide className="w-5 h-5" />
                                        {/if}
                                    </button>
                                </div>
                            </Input>
                            {#if form?.errors?.password}
                                <p>{form.errors.password}</p>
                            {/if}
                            <Input type={formDataObject.password_confirm.type} name={formDataObject.password_confirm.name} label={formDataObject.password_confirm.label} bind:value={formDataObject.password_confirm.value} additionalClasses="w-full" passwordClasses="flex flex-row justify-between items-center" />
                            {#if form?.errors?.passwordconfirm}
                                <p>{form.errors.passwordconfirm}</p>
                            {/if}
                    </div>
                    <div class="w-full md:flex md:justify-between md:items-center md:bg-white mt-4 md:px-2 md:rounded-md md:border md:border-zinc-200/[.7] md:py-3">
                        <div class="max-md:hidden">
                            <AuthSectionToggler />
                        </div>
                        <Button additionalClasses="w-full ">
                            Continuer
                        </Button>
                    </div>
                </form>
        {:else}
        <SearchCountry data={data} />
                <form action="?/register" method="POST" class="flex items-center justify-center flex-col" use:enhance={handleSubmit}>
                    <input type="hidden" name="step" value=1>
                    <div class="flex flex-col sm:flex-row sm:gap-x-1.5 w-full">
                        <div class="w-full">
                            <Input type={formDataObject.firstName.type} name={formDataObject.firstName.name} label={formDataObject.firstName.label} bind:value={formDataObject.firstName.value} additionalClasses="w-full" />
                            {#if form?.errors?.firstname}
                            <p class="text-red-400 mt-1">{form.errors.firstname}</p>
                            {/if}
                        </div>
                        <div class="w-full">
                            <Input type={formDataObject.lastName.type} name={formDataObject.lastName.name} label={formDataObject.lastName.label} bind:value={formDataObject.lastName.value} additionalClasses="w-full" />
                            {#if form?.errors?.lastname}
                            <p class="text-red-400 mt-1">{form.errors.lastname}</p>
                            {/if}
                        </div>
                    </div>
                    <div class="w-full">
                        <Input type={formDataObject.email.type} name={formDataObject.email.name} label={formDataObject.email.label} bind:value={formDataObject.email.value} additionalClasses="w-full"  />
                        {#if form?.errors?.email}
                            <p class="text-red-400 mt-1">{form?.errors?.email}</p>
                        {/if}
                    </div>
                        <input type="hidden" name="phone_number_country_prefix" value={$countryStore.code}>
                        <div class="w-full flex flex-col items-center">
                            <Input type={formDataObject.phone_number.type} name={formDataObject.phone_number.name} label={formDataObject.phone_number.label} bind:value={formDataObject.phone_number.value} additionalClasses="w-full" />
                            {#if form?.errors?.phone_number}
                                <p class="text-red-400 mt-1 w-full">{form?.errors?.phone_number}</p>
                            {/if}
                        </div>
                        <div class="w-full md:flex md:justify-between md:items-center md:bg-white mt-4 md:px-2 md:rounded-md md:border md:border-zinc-200/[.7] md:py-3">
                            <div class="max-md:hidden">
                                <AuthSectionToggler />
                            </div>
                            <Button additionalClasses="w-full ">
                                Continuer
                            </Button>
                        </div>
                </form>
                {/if}
            </div>
            <div class="mt-4 md:hidden">
                <AuthSectionToggler />
            </div>
        </div>