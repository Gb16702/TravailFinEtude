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
    import {isOpen} from "../../../stores/openStore"

    console.log($isOpen);

    export const snapshot: Snapshot = {
        capture: () => formDataObject,
        restore: data => formDataObject = data
    }

    const togglePasswordVisibility: () => "text" | "password" = () => {
        return formDataObject.password.type = formDataObject.password_confirm.type = formDataObject.password.type === "password" ? "text" : "password"
    };

    let resultValue: any = null;
    let step: any = null

    const handleSubmit:SubmitFunction = () => {
        return async ({result, update}) => {
            resultValue = result;
            step = resultValue?.data?.step
            await update();
        }
    }
    export let form:ActionData;

    const totalSteps = 3;

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
            name:"firstname",
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
            placeholder: "Votre numéro de téléphone",
            label: "Numéro de téléphone",
            value: "",
        },
        password: {
            type: "password",
            name:"password",
            placeholder: "•••••••",
            label: "Mot de passe",
            value: "",
        },
        password_confirm: {
            type: "password",
            name:"passwordconfirm",
            placeholder: "•••••••",
            label: "Confirmer le mot de passe",
            value: "",
        }
    }

    let checked = true;

    export let data: PageData

</script>

    <div class="w-full flex items-center justify-center flex-col bg-white h-[70%]">
        <div class="w-[97%] md:w-[75%] lg:w-[650px] px-2 flex flex-col gap-y-5 justify-center">
            <div class="w-full flex items-start justify-center flex-col gap-y-1 mb-[12px]">
                <h1 class="font-bold text-[28px] text-accent-violet-d">Inscription</h1>
                <h3 class="text-[15px] text-zinc-300">{`Etape ${step ?? 1} sur ${totalSteps}`}</h3>
            </div>
        {#if step === 3}
            <div>
                <form action="?/register" class="rounded-[5px] overflow-hidden" method="POST" use:enhance={handleSubmit}>
                    <input type="hidden" name="step" value=3>
                        <div class="flex flex-col ">
                        <section class=" pb-7 pt-1">
                            <h2 class="text-left font-semibold text-[16.5px]">Récapituatif de l'inscription</h2>
                            <div class="border-l-2 border-zinc-100 pl-3">
                                <div class="flex flex-row w-full gap-x-2 mt-7 h-[35px]">
                                    <label for="firstName" class="flex flex-col w-[35%] h-full font-medium text-sm">
                                        Prénom
                                    </label>
                                    <input
                                        class="px-2 border border-zinc-200/[.5] overflow-visible rounded-md w-[65%]"
                                        name={formDataObject.firstName.name}
                                        placeholder={formDataObject.firstName.placeholder}
                                        type={formDataObject.firstName.type}
                                        value={formDataObject.firstName.value}
                                    />
                                </div>
                                <div class="flex flex-row w-full gap-x-2 mt-4 h-[35px]">
                                    <label for="firstName" class="flex flex-col w-[35%] h-full font-medium text-sm">
                                        Nom
                                    </label>
                                    <input
                                    class="px-2 border border-zinc-200/[.5] rounded-md w-[65%]"
                                    name={formDataObject.lastName.name}
                                    placeholder={formDataObject.lastName.placeholder}
                                    type={formDataObject.lastName.type}
                                    value={formDataObject.lastName.value}
                                    />
                                </div>
                                <div class="flex flex-row w-full gap-x-2 mt-4 h-[35px]">
                                    <label for="firstName" class="flex flex-col w-[35%] h-full">
                                        Adresse mail
                                    </label>
                                    <input
                                    class="px-2 border border-zinc-200/[.5] rounded-md w-[65%]"
                                    name={formDataObject.email.name}
                                    placeholder={formDataObject.email.placeholder}
                                    type={formDataObject.email.type}
                                    value={formDataObject.email.value}
                                    />
                                </div>
                                <div class="flex flex-row w-full gap-x-2 mt-4 h-[35px]">
                                    <label for="firstName" class="flex flex-col w-[35%] h-full">
                                        Mot de passe
                                    </label>
                                    <input
                                    class="px-2 border border-zinc-200/[.5] rounded-md w-[65%]"
                                    name={formDataObject.password.name}
                                    placeholder={formDataObject.password.placeholder}
                                    type={formDataObject.password.type}
                                    value={formDataObject.password.value}
                                    />
                                </div>
                            </div>
                        </section>
                        <section class="flex row w-full flex-1 bg-zinc-100/[.8] min-h-[70px] rounded-[5px]">
                            <div class=" w-[10%] flex items-center justify-center">
                                <Checkbox />
                            </div>
                            <div class="w-[90%] flex items-center justify-center] pr-1">
                                <small class="text-[14px] text-zinc-600">
                                    En cochant cette case, vous acceptez les
                                    <span class="text-accent-green font-medium">
                                        Conditions Générales d'Utilisation
                                    </span>
                                </small>
                            </div>
                        </section>
                    </div>
                    <div class="w-full h-[80px] flex flex-row items-center justify-between">
                        <Button additionalClasses="w-full">
                            Terminer
                        </Button>
                    </div>
                </form>
            </div>
        {:else if step === 2}
                    <form action="?/register" method="POST" use:enhance={handleSubmit}>
                        <input type="hidden" name="step" value=2>
                        <div class="flex flex-col gap-y-4 sm:flex-row sm:gap-x-1.5 w-full">
                            <div>
                                <Input placeholder={formDataObject.password.placeholder} type={formDataObject.password.type} name={formDataObject.password.name} label={formDataObject.password.label} bind:value={formDataObject.password.value} additionalClasses="w-full"
                                    passwordClasses="flex flex-row justify-between items-center">
                                    <SearchCountry data={data} />
                                    <button on:click|preventDefault={togglePasswordVisibility}>
                                        {#if formDataObject.password.type==="password"}
                                            <Show className="w-5 h-5 stroke-zinc-500    " />
                                        {:else if formDataObject.password.type==="text"}
                                            <Hide className="w-5 h-5" />
                                        {/if}
                                    </button>
                                </Input>
                                {#if form?.errors?.password}
                                    <p>{form.errors.password}</p>
                                {/if}
                            </div>
                            <div>
                                <Input placeholder={formDataObject.password_confirm.placeholder} type={formDataObject.password_confirm.type} name={formDataObject.password_confirm.name} label={formDataObject.password_confirm.label} bind:value={formDataObject.password_confirm.value} additionalClasses="w-full" passwordClasses="flex flex-row justify-between items-center" />

                                {#if form?.errors?.password_confirm}
                                    <p>{form.errors.password_confirm}</p>
                                {/if}
                            </div>
                        </div>
                        <Button additionalClasses="w-full sm:w-[340px] mt-4">
                        Continuer
                        </Button>
                    </form>
        {:else}
        <SearchCountry data={data} />
            <form action="?/register" method="POST" class="flex items-center justify-center flex-col" use:enhance={handleSubmit}>
                <input type="hidden" name="step" value=1>
                    <div class="flex flex-col gap-y-4 sm:flex-row sm:gap-x-1.5 w-full">
                            <div>
                            <Input placeholder={formDataObject.firstName.placeholder} type={formDataObject.firstName.type} name={formDataObject.firstName.name} label={formDataObject.firstName.label} bind:value={formDataObject.firstName.value} additionalClasses="w-full" />
                                {#if form?.errors?.firstname}
                                    <p class="text-red-400 mt-1">{form.errors.firstname}</p>
                                {/if}
                            </div>
                            <div>
                            <Input placeholder={formDataObject.lastName.placeholder} type={formDataObject.lastName.type} name={formDataObject.lastName.name} label={formDataObject.lastName.label} bind:value={formDataObject.lastName.value} additionalClasses="w-full" />
                                {#if form?.errors?.lastname}
                                    <p class="text-red-400 mt-1">{form.errors.lastname}</p>
                                {/if}
                            </div>
                    </div>
                    <div class="mt-4 w-full">
                        <Input placeholder={formDataObject.email.placeholder} type={formDataObject.email.type} name={formDataObject.email.name} label={formDataObject.email.label} bind:value={formDataObject.email.value} additionalClasses="w-full"  />
                    {#if form?.errors?.email}
                        <p class="text-red-400 mt-1">{form?.errors?.email}</p>
                    {/if}
                    </div>
                    <div class="mt-4 w-full flex flex-rox items-center">
                        <Input placeholder={formDataObject.phone_number.placeholder} type={formDataObject.phone_number.type} name={formDataObject.phone_number.name} label={formDataObject.phone_number.label} bind:value={formDataObject.phone_number.value} additionalClasses="w-full" />
                        {#if form?.errors?.phone_number}
                            <p class="text-red-400 mt-1">{form?.errors?.phone_number}</p>
                        {/if}
                    </div>
                    <Button additionalClasses="w-full sm:w-[340px] mt-4">
                        Continuer
                    </Button>
                </form>
            <AuthSectionToggler />
        {/if}
        </div>
    </div>
