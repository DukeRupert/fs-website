<script lang="ts">
  import PageContainer from "$lib/components/PageContainer.svelte";
  import PageIntro from "$lib/components/PageIntro.svelte";
  import ProjectShowcase from "../ProjectShowcase.svelte";
  import eandi_logo from "$lib/assets/images/ebony_and_ivory/logo.png?as=run";
  import eandi_chante from "$lib/assets/images/ebony_and_ivory/chante.png?as=run";
  import eandi_splash from "$lib/assets/images/ebony_and_ivory/splash_mobile.png?as=run";

  const caseStudy = {
    client: "Ebony and Ivory",
    title: "A passion for music in Montana",
    description:
      "In the charming town of Helena, MT, there lived a music virtuoso whose unwavering passion for both music and children led her to open her very own enchanted studio for the sweet melodies of young hearts to flourish.",
    href: "/portfolio/ebony-and-ivory",
    link: "https://www.ebonyandivory.love/",
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
    splash: {
      src: eandi_splash,
      alt: "A mobile view of Ebony and Ivory website",
    },
  };

  const pages = [
    {
      href: "/portfolio/kagen-coffee-and-crepes",
      client: "Kagen Coffee and Crepes",
      title: "Crepes, coffee, and community",
      description:
        "A vacation led to love, which led to a dream, which became Kagen Coffee and Crepes in Richland, WA",
      date: "2022-06",
    },
    {
      href: "/portfolio/fts-excavation",
      client: "FtS Excavation",
      title: "Precision + Integrity",
      description:
        "A dedication to bring precision and integrity to every excavation job in the challenging Puget Sound region.",

      date: "2021-08",
    },
  ];
</script>

<svelte:head>
  <title>Firefly Software | Our Work | Ebony and Ivory</title>
</svelte:head>

<PageContainer centered={true}>
  <span slot="intro">
    <PageIntro eyebrow="Case Study" title={caseStudy.title} centered={true}>
      <p>
        {caseStudy.description}
      </p>
    </PageIntro>
  </span>

  <span slot="body">
    <div class="mx-auto max-w-5xl">
      <ProjectShowcase {caseStudy} {pages} />
    </div>
  </span>
</PageContainer>
