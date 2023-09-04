import fts_logo from "$lib/assets/images/fts_logo.png?as=run";
import fts_frank from "$lib/assets/images/fts_frank.jpg?as=run";
import kcc_logo from "$lib/assets/images/kcc_logo.png?as=run";
import kcc_kagen from "$lib/assets/images/kagen_coffee_and_crepes_owner_873.webp?as=run";
import eandi_logo from "$lib/assets/images/ebony_and_ivory/logo.png?as=run:0";
import eandi_chante from "$lib/assets/images/ebony_and_ivory/chante.png?as=run";

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
    client: "Ebony and Ivory",
    title: "A passion for music in Montana",
    description:
      "In the charming town of Helena, MT, there lived a music virtuoso whose unwavering passion for both music and children led her to open her very own enchanted studio for the sweet melodies of young hearts to flourish.",
    href: "/portfolio/ebony-and-ivory",
    summary: [
      "Working alongside the Ebony and Ivory music lessons team to bring their new website to life was a fantastic experience. From the initial discussions, it was clear that they had a clear vision and purpose for their online presence, which made the entire development process seamless and enjoyable. The end result is a visually stunning and highly user-friendly website that aligns perfectly with their goals. It's been incredibly satisfying to see the website's effectiveness in generating a consistent stream of new clients for their business.",
      "Since the website's launch, we've observed a significant increase in inquiries and bookings, which validates our approach to design and functionality. Our attention to detail and dedication to Ebony and Ivory's success have been key factors in achieving these results. It's rewarding to witness their music lessons reaching a broader audience, and we're excited to continue supporting the growth of Ebony and Ivory.",
    ],
    logo: {
      src: eandi_logo,
      alt: "Ebony and Ivory logo",
    },
    skills: ["Web Design", "Web Development", "Hosting", "Social Media"],
    date: "2023-07",
    service: "Website development",
    testimonial: {
      author: { name: "Chante' Williams", role: "Owner" },
      src: eandi_chante,
      content: `As a business owner, I'm all about efficiency, and Firefly Software totally delivered. Thanks to their work, our online presence has blown up, and we've got more students than ever.`,
    },
  },
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
