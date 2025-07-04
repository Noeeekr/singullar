// Components
import Stack from "@mui/material/Stack";
import Typography from "@mui/material/Typography";
import CircularButton from "@components/ButtonCircular";

// Features
import { useNavigate } from "react-router-dom";

/**
 * Contains the title, a button to navigate back and a text subtitle that is a breadcrumbs by default.
 *
 * Loads children in right side.
 */
const SectionHeader = ({
  children,
  title,
  subtitle,
  breadcrumbs,
}: {
  children?: JSX.Element;
  breadcrumbs?: JSX.Element;
  title: string;
  subtitle?: string;
}): JSX.Element => {
  const navigate = useNavigate();
  
  return (
    <Stack direction="row" alignItems="center" width="100%" gap={1}>
      <CircularButton
        onClickCb={() => {
          navigate(-1);
        }}
      />
      <Stack gap={0.5}>
        <Typography component="h4" variant="h4" fontWeight="600">
          {title}
        </Typography>
        <Stack direction="row" gap={0.5}>
            {
              subtitle 
              ? <Typography fontWeight={500} sx={{ color: 'rgb(130,130,130)'}}> 
                { subtitle }
              </Typography>
              : breadcrumbs
            }
        </Stack>
      </Stack>
      {children || <></>}
    </Stack>
  );
};

export default SectionHeader;
