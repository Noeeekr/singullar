import Box from "@mui/material/Box"
import Typography from "@mui/material/Typography"
import ErrorBubble from "@components/bubbles/ErrorBubble/ErrorBubble"
import Stack, { StackProps } from "@mui/material/Stack"
import Grid from "@mui/material/Grid2"
import Divider from "@mui/material/Divider"

import useFetchClasses from "@hooks/useFetchClasses"
import { useEffect } from "react"

import { useParams } from "react-router-dom"
import SectionHeader from "@components/headers/sectionHeader/SectionHeader"
import useFetchStudents from "@hooks/useFetchStudents"
import useFetchUsers from "@hooks/useFetchUsers"
import { ROLE_TEACHER } from "../../../../models/server"

const ClassPage = ({ ...props }: StackProps): JSX.Element => {
    const classRequest = useFetchClasses()
    const studentRequest = useFetchStudents()
    const teacherRequest = useFetchUsers()

    const params = useParams()
    const classe = classRequest.response?.length ? classRequest.response[0] : null
    

    useEffect(() => {
        classRequest.send([{ id: Number(params["id"]) }])
    }, [])

    useEffect(() => {
        if (classe == null) return; 
        studentRequest.send([
            { class_id: classe.id, segment: classe.segment }
        ])
        teacherRequest.send({ 
            "accepted_roles": [ROLE_TEACHER],
            "offset": 0,
            "filters": [
                { "id": classe.teacher_id },
            ]
        })
    }, [classe])
 
    if (classRequest.error) return <ErrorBubble message={classRequest.error} />
    if (classRequest.isLoading) return <Typography>Carregando Informações</Typography>
    if (classe == null) return <ErrorBubble message="Nenhuma turma encontrada" />

    return (
        <Stack {...props} gap={2}>
            <SectionHeader
                title="Informações gerais"
                subtitle="Analise e verifique as informações da turma"
            />
            <Stack gap={2}>
                <Divider/>
                <Stack>
                    <Typography variant="h5" component="p" textTransform="capitalize"><span style={{ fontWeight: "bold" }}>Plataforma ID da classe:</span> {classe.id + 1_000_000}</Typography>
                    <Typography variant="h5" component="p" textTransform="capitalize"><span style={{ fontWeight: "bold" }}>Nome da classe:</span> {classe.name}</Typography>
                    <Typography variant="h5" component="p" textTransform="capitalize"><span style={{ fontWeight: "bold" }}>Segmento da classe:</span> {classe.segment}</Typography>
                    <Typography variant="h5" component="p" textTransform="capitalize"><span style={{ fontWeight: "bold" }}>Série da classe:</span> {classe.series}</Typography>
                </Stack>
                <Divider/>
                <Typography variant="h5" component="p" textTransform="capitalize">Informações dos estudantes:</Typography>
                {
                    studentRequest.error
                        ? <ErrorBubble message={studentRequest.error} />
                        : <></>
                }
                {
                    studentRequest.isLoading
                        ? <Typography fontWeight="bold" variant="body1" component="p">Carregando Estudantes</Typography>
                        : <></>
                }
                {
                    studentRequest.response?.length
                        ? <Grid container gap={1}>
                            {
                                studentRequest.response.map((student) => (
                                    <Grid size={5}>
                                        <Box sx={{
                                            backgroundColor: "white", 
                                            paddingY: 1.5,
                                            paddingX: 4,
                                            borderRadius: 4,
                                            border: "solid 1px rgb(230,230,230)",
                                            boxShadow: "0px 0px 2px 3px rgb(100,100,100,0.05)",
                                        }}>
                                            <Typography fontWeight="bold">{student.name}</Typography>
                                            <Typography>{student.email}</Typography>
                                            <Typography textTransform="capitalize">{student.segment}</Typography>
                                        </Box>
                                    </Grid>
                                ))
                            }
                        </Grid>
                        : <ErrorBubble message={"Nenhum estudante encontrado"} />
                }
                <Divider/>
                <Typography variant="h5" component="p" textTransform="capitalize">Informações dos professores:</Typography>
                {
                    studentRequest.error
                        ? <ErrorBubble message={studentRequest.error} />
                        : <></>
                }
                {
                    studentRequest.isLoading
                        ? <Typography fontWeight="bold" variant="body1" component="p">Carregando Estudantes</Typography>
                        : <></>
                }
                {
                    teacherRequest.response?.length
                        ? <Grid container gap={1}>
                            {
                                teacherRequest.response.map((teacher) => (
                                    <Grid size={5}>
                                        <Box sx={{
                                            backgroundColor: "white", 
                                            paddingY: 1.5,
                                            paddingX: 4,
                                            borderRadius: 4,
                                            border: "solid 1px rgb(230,230,230)",
                                            boxShadow: "0px 0px 2px 3px rgb(100,100,100,0.05)",
                                        }}>
                                            <Typography fontWeight="bold">{teacher.name}</Typography>
                                            <Typography>{teacher.email}</Typography>
                                        </Box>
                                    </Grid>
                                ))
                            }
                        </Grid>
                        : <ErrorBubble message={"Nenhum professor encontrado"} />
                }
            </Stack>
        </Stack>
    )
}

export default ClassPage;