// Components
import Typography from '@mui/material/Typography';
import Stack from '@mui/material/Stack';
import Box from '@mui/material/Box';

import ButtonSolid from '@components/ButtonSolid';

// Features
import { useState, createContext, useEffect, Fragment } from 'react';
import { useForm } from 'react-hook-form';

// Types
import type { IFirstSectionData } from './CreateFirstSection';
import type { ISecondSectionData } from './CreateSecondSection';
import type { IThirdSectionData } from './CreateThirdSection';

import type { UseFormSetValue } from 'react-hook-form';

export const CreateFormContext = createContext<{
    setFormValue: UseFormSetValue<IFirstSectionData & ISecondSectionData & IThirdSectionData>,
    setSectionStatus: Function,

    // Enables react hook form to take control of the input via parent Controller component;
    formState: IFirstSectionData & ISecondSectionData & IThirdSectionData,
}>({
    setSectionStatus: () => { console.log("If this is showing, it means a context function that's being used is not implemented in some piece of code. This might end in code being executed incorrectly. Please spend some time debugging this page component") },
    formState: {} as IFirstSectionData & ISecondSectionData & IThirdSectionData,
    setFormValue: {} as UseFormSetValue<IFirstSectionData & ISecondSectionData & IThirdSectionData>,
});

const CreateClassForm = ({ form }: { form: ({ title: string, subtitle: string, content: JSX.Element })[]}): JSX.Element => {
    const { setValue, handleSubmit, register, watch } = useForm<IFirstSectionData & ISecondSectionData & IThirdSectionData>({
        defaultValues: {
            currentYear: -1, // PT: Ano letivo
            segment: -1, // PT: segmento
            series: -1, // PT: série/ano
            studentSheet: -1, // For adding students ; PT: Planilha CSV
        }
    });
    const currentYear = watch("currentYear");
    const segment = watch("segment");
    const series = watch("series");

    // define what stage of form is open
    let [sectionsStatus, setSectionStatus] = useState<1 | 2 | 3>(1)
    

    const onSubmit = () => {

    };

    useEffect(() => {
        register('currentYear')
        register('segment')
        register('series')
    },[register])

    // Handle next section 
    useEffect(() => {
        if (series != -1 && segment != -1 && currentYear != -1) {
            setSectionStatus(2)
        }
    },[series, segment, currentYear])

    return(
        <form onSubmit={handleSubmit(onSubmit)}>
                <CreateFormContext.Provider value={{
                    setSectionStatus: setSectionStatus,
                    setFormValue: setValue,
                    formState: { currentYear, segment, series, studentSheet: -1 },
                }}>
                    {
                        // MAP OVER FORM SECTIONS DATA TO TURN INTO JSX;
                        form.map((data, i) =>
                            <Fragment key={data.title}>
                                <Stack marginBottom={4}>
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
                                            // REDO BUTTON ;
                                            i + 1 > 1 && sectionsStatus === i + 1 && <div style={{ margin: '0 0 0 auto'}}>
                                                <ButtonSolid 
                                                    onClick={() => {
                                                        if (sectionsStatus == 2) {
                                                            setValue("currentYear",-1)
                                                            setValue("segment", -1)
                                                            setValue("series",-1)
                                                        }
                                                        if (sectionsStatus == 3) {
                                                            setValue("studentSheet",-1)
                                                        }
                                                        setSectionStatus(prev => (prev - 1) as 1 | 2 );
                                                    }}
                                                >
                                                    Refazer última etápa
                                                </ButtonSolid>
                                            </div>
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
                            </Fragment>
                        )
                    }
                </CreateFormContext.Provider>
            </form>
    )
}

export default CreateClassForm;