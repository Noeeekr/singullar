// Components
import OutlinedInput from "@mui/material/OutlinedInput";
import FormControl from "@mui/material/FormControl";
import InputLabel from "@mui/material/InputLabel";
import Stack from "@mui/material/Stack";
import SectionHeader from "@components/headers/sectionHeader/SectionHeader";
import Grid from "@mui/material/Grid2";
import ButtonSolid from "@components/buttons/Default/Solid";
import ErrorHelperText from "@components/bubbles/ErrorBubble/ErrorHelperText";

// Features
import { useRef } from "react";
import { useForm } from "react-hook-form";

// Types
import { SubmitHandler } from "react-hook-form";

interface INewbornStudent {
  email: string;
  name: string;
}

interface IReqCreateStudents {
  students: {
    [index: string]: string;
  };
  institutionId: number;
}

interface IFormNecessities extends INewbornStudent {
  emailsInUse: {
    [index: string]: boolean;
  };
}
const CreateStudent = (): JSX.Element => {
  const {
    handleSubmit,
    register,
    watch,
    setError,
    setValue,
    getValues,
    formState: { errors },
  } = useForm<IReqCreateStudents & IFormNecessities>({
    defaultValues: {
      name: "",
      email: "",
      institutionId: 1,
      students: {},
      emailsInUse: {},
    },
  });

  const secondInput = useRef<HTMLInputElement>();
  const submitButton = useRef<HTMLInputElement | null>(null);

  const students = watch("students");

  const onSubmit: SubmitHandler<INewbornStudent> = (data) => {
    const stds = getValues().students;
    const emailsInUse = getValues().emailsInUse;

    if (stds[data.name.trim()]) {
      // Get older email to remove it from emails already in use
      const oldEmail = stds[data.name.trim()];

      emailsInUse[oldEmail] = false;
    } else if (emailsInUse[data.email]) {
      // Check if e-mail is in use and throws err
      setError("email", {
        message:
          "E-mail já selecionado. Para continuar, mude o e-mail do estudante. ",
      });
      return;
    }

    stds[data.name.trim()] = data.email;
    emailsInUse[data.email] = true;

    setValue("emailsInUse", emailsInUse);
    setValue("students", stds);
  };

  return (
    <div>
      <form onSubmit={handleSubmit(onSubmit)}>
        <SectionHeader
          title="Adicionar estudantes"
          subtitle="Defina as informações necessárias para adicionar os estudantes"
        >
          <ButtonSolid title="Enviar Formulário" sx={{ margin: "0 0 0 auto" }} disabled={true} />
        </SectionHeader>
        <Stack gap={2} direction="row" sx={{ alignItems: "center" }}>
          <Grid
            container
            spacing={2}
            sx={{ width: "100%", margin: "20px 0px" }}
          >
            <Grid size={{ mobile: 12, xss: 6 }}>
              <FormControl>
                <InputLabel htmlFor="student-create-email-input">
                  Nome completo do estudante
                </InputLabel>
                <OutlinedInput
                  {...register("name", {
                    required: "Por favor insira um nome.",
                    pattern: {
                      value: /^[a-zA-Záàãâäéè êëíìîïóòõôöúùûüçÇ]+$/g,
                      message: "Por favor insira um nome válido",
                    },
                    minLength: {
                      value: 5,
                      message: "O nome deve ter no minímo cinco caracteres.",
                    },
                  })}
                  error={Boolean(errors?.name)}
                  id="student-create-email-input"
                  label="nome-completo-do-estudante"
                  onKeyDown={(e) => {
                    if (e.key === "Enter") {
                      secondInput?.current?.focus();
                    }
                  }}
                />
                <ErrorHelperText show={Boolean(errors?.name)}>
                  {errors?.name?.message}
                </ErrorHelperText>
              </FormControl>
            </Grid>
            <Grid size={{ mobile: 12, xss: 6 }}>
              <FormControl>
                <InputLabel htmlFor="student-create-plataformId-input">
                  Email do estudante
                </InputLabel>
                <OutlinedInput
                  {...register("email", {
                    required: "Por favor insira um e-mail.",
                    pattern: {
                      value:
                        /^(([^<>()[\]\\.,;:\s@"]+(\.[^<>()[\]\\.,;:\s@"]+)*)|(".+"))@((\[[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}])|(([a-zA-Z\-0-9]+\.)+[a-zA-Z]{2,}))$/,
                      message: "Por favor insira um email válido.",
                    },
                  })}
                  error={Boolean(errors?.email)}
                  inputRef={secondInput}
                  label="email-do-estudante"
                  id="student-create-plataformId-input"
                  onKeyDown={(e) => {
                    if (e.key === "Enter") {
                      submitButton?.current?.click();
                    }
                  }}
                />
                <ErrorHelperText show={Boolean(errors?.email)}>
                  {errors?.email?.message}
                </ErrorHelperText>
              </FormControl>
            </Grid>
            <Grid
              size={{ mobile: 12, xss: 6 }}
              sx={{ margin: "0 0 auto auto", maxWidth: "250px" }}
            >
              <ButtonSolid
                title="Adicionar Estudante"
                onClick={() => {
                  submitButton.current?.click();
                }}
              />
            </Grid>
          </Grid>
        </Stack>
        <input
          ref={submitButton}
          type="submit"
          style={{ pointerEvents: "none", opacity: 0 }}
        />
      </form>
      <div style={{ display: "flex", flexDirection: "column", gap: "10px" }}>
        {Object.entries(students).map((std_info) => (
          <div style={{ minHeight: '70px', width: '100%', display: 'flex', boxShadow: '0px  5px 5px 1px rgb(0,0,0,0.1)' }}>
            <div
              style={{
                width: '100%',
                gap: "10px",
                padding: "10px 7px",
                display: "flex",
                justifyContent: 'flex-between',
                alignItems: "center",
                backgroundColor: 'white',
              }}
            >
              <p style={{ width: '50%', textWrap: 'nowrap', textOverflow: "ellipsis", margin: 'auto 10px auto auto', boxSizing: 'border-box', overflow: 'hidden' }}>{std_info[0]}</p>
              <p style={{ width: '50%', textOverflow: "ellipsis", margin: 'auto 10px auto auto', boxSizing: 'border-box', overflow: 'hidden' }}>{std_info[1]}</p>
            </div>
            <div style={{
              display: 'flex',
              alignItems: 'center',
              justifyContent: 'center',
              backgroundColor: "rgba(255,100,100,1)",
              borderRadius: '0px 10px 10px 0px'
            }}>
              <p
                style={{
                  cursor: "pointer",
                  color: "rgba(255,255,255,0.8)",
                  margin: "0px 0px 0px 10px",
                  padding: "5px 12px",
                  fontWeight: "bold",
                }}
              >
                Excluir
              </p>
            </div>
          </div>
        ))}
      </div>
    </div>
  );
};

export default CreateStudent;
