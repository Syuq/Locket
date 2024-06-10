import { Link } from "@mui/joy";
import Icon from "@/components/Icon";
import MobileHeader from "@/components/MobileHeader";

const About = () => {
  return (
    <section className="@container w-full max-w-5xl min-h-full flex flex-col justify-start items-center sm:pt-3 md:pt-6 pb-8">
      <MobileHeader />
      <div className="w-full px-4 sm:px-6">
        <div className="w-full shadow flex flex-col justify-start items-start px-4 py-3 rounded-xl bg-white dark:bg-zinc-800 text-black dark:text-gray-300">
          <a href="https://www.duyquys.id.vn" target="_blank">
            <img className="w-auto h-12" src="https://www.duyquys.id.vn/full-logo-landscape.png" alt="locket" />
          </a>
          <p className="text-base">LOCKET IMAGE SHARING APP AND MORE.</p>
          <div className="mt-1 flex flex-row items-center flex-wrap">
            <Link underline="always" href="https://www.github.com/Syuq/Locket" target="_blank">
              GitHub Repo
            </Link>
            <Icon.Dot className="w-4 h-auto opacity-60" />
            <Link underline="always" href="https://www.duyquys.id.vn/" target="_blank">
              Official Website
            </Link>
          </div>
        </div>
      </div>
    </section>
  );
};

export default About;
