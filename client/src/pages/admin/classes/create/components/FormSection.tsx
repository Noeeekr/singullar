import Stack from "@mui/material/Stack"
import Typography from "@mui/material/Typography"
import Box from "@mui/material/Box"
import ButtonSolid from "@components/ButtonSolid"

import { FormContext } from "./Form"
import { formSections as sections } from "../data"

import { useCallback, useContext, useMemo } from "react"

import type { FormRequest, FormSections } from "./Form"

export default function FormSection({ activeSection }: { activeSection: number }): JSX.Element {
    const { setActiveSection, setFormData, formSections, send } = useContext(FormContext);

    const formData = useMemo(() => {
        let data: FormRequest = {} as FormRequest
    
        const keys = Object.keys(formSections) as Array<keyof FormSections>
        for (const key of keys) {
            const section = formSections[key]
            data = { ...data, ...section }
        }
        return data
    }, [formSections])

    const handleRedoSection = useCallback(() => {
        setActiveSection((section) => {
            if (section == 1) {
                return 1
            }
            switch(section) {
                case 2:
                    setFormData("firstSection.className", "")
                    setFormData("firstSection.segment", "")
                    setFormData("firstSection.series", "")
                break
                case 3:
                    setFormData("secondSection.students", [])
                break
                case 4:
                    setFormData("thirdSection.teacher_id", null)
                break
            }
            return section - 1 as typeof section
        });
    }, [setActiveSection, setFormData])
    return (
        <div>
            {
                sections.map((section, i) =>
                    <Stack key={section.title} marginBottom={4}>
                        <Stack direction="row" justifyContent="space-between">
                            <Stack direction="column">
                                <Typography variant="h5" component="h4" color="primary.purpleDark" fontWeight={600}>
                                    {section.title}
                                </Typography>
                                {
                                    // Subtitle if exists
                                    activeSection === i + 1
                                        ?
                                        <Typography variant="body1" component="p" color="rgb(145,145,145)" fontWeight={600}>
                                            {section.subtitle}
                                        </Typography>
                                        : <></>
                                }
                            </Stack>
                            {
                                // Redo Button + Custom Button
                                activeSection !== (i + 1)
                                    ? <></>
                                    : <div style={{ margin: '0 0 0 auto' }}>
                                        <Stack flexDirection="row" gap={2}>
                                            {
                                                i == 1
                                                    ? <></>
                                                    : <ButtonSolid onClick={handleRedoSection}>
                                                        Refazer última etápa
                                                    </ButtonSolid>
                                            }
                                            {
                                                section.button != null
                                                    ? section.button
                                                    : activeSection == 4 
                                                    ? <ButtonSolid onClick={() => { send(formData); setActiveSection(1) }}>
                                                        Criar turma
                                                    </ButtonSolid>
                                                    : <ButtonSolid onClick={() => {setActiveSection((section) => section + 1 as typeof section)}}>
                                                        Continuar
                                                    </ButtonSolid>
                                            }
                                        </Stack>
                                    </div>
                            }
                        </Stack>
                        <Box
                            height={activeSection == i + 1 ? "100%" : 0}
                            display={activeSection == i + 1 ? "initial" : "none"}
                            marginTop={0.5}
                        >
                            {section.content}
                        </Box>
                    </Stack>
                )
            }
        </div>
    )
}