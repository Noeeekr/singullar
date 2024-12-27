// Components
import Typography from '@mui/material/Typography';
import Stack from '@mui/material/Stack';
import Box from '@mui/material/Box';

import ButtonSolid from '@components/ButtonSolid';

// Features
import { useState, createContext, useContext, useEffect, Fragment } from 'react';
import { useForm } from 'react-hook-form';

// Types
import type { IFirstSectionData } from './CreateFirstSection';

import type { UseFormSetValue } from 'react-hook-form';

import type { Dispatch, SetStateAction } from 'react';

export const CreateFormContext = createContext<{
    setFormValue: UseFormSetValue<IFirstSectionData>,
    setSectionStatus: Dispatch<SetStateAction<1 | 2 | 3>>,

    // Enables react hook form to take control of the input via parent Controller component;
    formState: IFirstSectionData,
}>({
    setSectionStatus: () => { console.log("If this is showing, it means a context function that's being used is not implemented in some piece of code. This might end in code being executed incorrectly. Please spend some time debugging this page component") },
    formState: {} as IFirstSectionData,
    setFormValue: {} as UseFormSetValue<IFirstSectionData>,
});


export const ButtonRedoLastSection = (): JSX.Element => {
    const { setFormValue, setSectionStatus, formState } = useContext(CreateFormContext);

    return(
        <div style={{ margin: '0 0 0 auto'}}>
            <ButtonSolid 
                onClick={() => {
                    setSectionStatus(prev => {
                        if (prev === 2) {
                            setFormValue("currentYear",-1)
                            setFormValue("segment", -1)
                            setFormValue("series",-1)

                            return 1     
                        }

                        if (prev === 3) {
                            setFormValue("class", { ...formState.class, students: [] })
                            return 2
                        }

                        return 3
                    });
                }}
            >
                Refazer última etápa
            </ButtonSolid>
        </div>
    )
}
/**
 * Form to create classes. 
 * 
 * This component also has a context provided by the same file that can help manage its inputs, their values, and their state.
 * @param param0 
 * @returns 
 */
const CreateClassForm = ({ form }: { form: ({ title: string, subtitle: string, content: JSX.Element, button: JSX.Element })[]}): JSX.Element => {
    const { setValue, handleSubmit, register, watch } = useForm<IFirstSectionData>({
        defaultValues: {
            currentYear: -1, // PT: Ano letivo
            segment: -1, // PT: segmento
            series: -1, // PT: série/ano
            class: {},
        }
    });
    const currentYear = watch("currentYear");
    const segment = watch("segment");
    const series = watch("series");
    const classOj = watch("class");

    // define what stage of form is open
    let [sectionsStatus, setSectionStatus] = useState<1 | 2 | 3>(1)

    const onSubmit = () => {

    };

    useEffect(() => {
        register('currentYear')
        register('segment')
        register('series')
    },[register])

    return(
        <form onSubmit={handleSubmit(onSubmit)}>
                <CreateFormContext.Provider value={{
                    setSectionStatus: setSectionStatus,
                    setFormValue: setValue,
                    formState: { currentYear, segment, series, class: classOj },
                }}>
                    {
                        // MAP OVER FORM SECTIONS DATA TO TURN INTO JSX;
                        form.map((data, i) =>
                                <Stack
                                    key={data.title}
                                    marginBottom={4}
                                >
                                    <Stack direction="row">

                                        { /* TITLE AND SUBTITLE */}
                                        <Stack direction="column">
                                            <Typography
                                                variant="h5"
                                                component="h4"
                                                color="primary.purpleDark"
                                                fontWeight={600}
                                            >
                                                {data.title}
                                            </Typography>

                                            {
                                                sectionsStatus === i + 1 && 
                                                <Typography
                                                    variant="body1"
                                                    component="p"
                                                    color="rgb(145,145,145)"
                                                    fontWeight={600}
                                                >
                                                    {data.subtitle}
                                                </Typography>
                                            }
                                        </Stack>
                                        
                                        {
                                            /* BUTTON */
                                            sectionsStatus === (i + 1) ? data.button : <></>
                                        }

                                    </Stack>
                                    { /* FORM SECTIONS GO HERE */}
                                    <Box sx={{
                                        height: sectionsStatus === i + 1 ? "100%" : 0,

                                        overflow: "hidden",

                                        marginTop: 0.5,
                                    }}>
                                        {data.content}
                                    </Box>
                                </Stack>
                        )
                    }
                </CreateFormContext.Provider>
            </form>
    )
}

export default CreateClassForm;