import fts_logo from "$lib/assets/images/fts_logo.png?run";
import fts_frank from "$lib/assets/images/fts_frank.jpg?run";
import kcc_logo from "$lib/assets/images/kcc_logo.png?run";
import kcc_kagen from "$lib/assets/images/kagen_coffee_and_crepes_owner_873.webp?run";

export interface Project {
  client: string;
  title: string;
  description: string;
  href: string;
  summary: string[];
  logo: {
    src: Object[];
    alt: string;
  };
  skills: string[];
  date: string;
  service: string;
  testimonial: {
    author: { name: string; role: string };
    src: Object[];
    content: string;
  };
}

const PROJECTS: Project[] = [
  {
    client: "Kagen Coffee and Crepes",
    title: "Crepes, coffee, and community",
    description:
      "A vacation led to love, which led to a dream, which became Kagen Coffee and Crepes in Richland, WA",
    href: "/portfolio/kagen-coffee-and-crepes",
    summary: [
      "Like many entrepreneurs Kagen started his business with a basic do-it-yourself website. After his first year though he wisely realized that his business needed a better online presence in order to compete in the Tri-Cities area.",
      "When Kagen Coffee and Crepes approached us to optimize their online presence we delivered a highly performant website that highlights their special combination of delicious crepes and friendly atmosphere. It must have worked since now Kagen Coffee and Crepes is about to open their second location!",
    ],
    logo: {
      src: kcc_logo,
      alt: "Kagen Coffee and Crepes logo",
    },
    skills: [
      "Web Design",
      "Web Development",
      "Hosting",
      "CMS (Sanity)",
      "Backend (Shopify)",
      "Integration (Toast)",
    ],
    date: "2022-06",
    service: "Website development",
    testimonial: {
      author: { name: "Kagen Cox", role: "Owner" },
      src: kcc_kagen,
      content: `Firefly is an absolutely must have company working for you in your
              corner! They are super approachable, helpful and friendly! If you
              are looking to take your company to the next level I highly
              recommend that you give Firefly a call!`,
    },
  },
  {
    client: "FtS Excavation",
    title: "Precision + Integrity",
    description:
      "A dedication to bring precision and integrity to every excavation job in the challenging Puget Sound region.",
    href: "/portfolio/fts-excavation",
    summary: [
      " Frank and Beverly founded FtS Excavation in 2019. The mission was to provide excavation services to the Puget Sound region in a manner that demonstrated the highest precision and integrity. We share their love of excellence and constant drive towards improvement.",
    ],
    logo: {
      src: fts_logo,
      alt: "FtS Excavation logo",
    },
    date: "2021-08",
    skills: ["Web design", "Web development"],
    service: "Website development",
    testimonial: {
      author: { name: "Frank Sharp", role: "Owner" },
      content: `Firefly is absolutely the best way to get the most value from your website. Reasonably priced and headache free! I appreciate getting to speak with a real person who responds lightening fast, listens to what I want and can create anything and do it better than anyone. I highly recommend using Firefly software to create your website or any digital tool you need. We've decided to hire them for our app creation too.`,
      src: fts_frank,
    },
  },
];

export default PROJECTS;
