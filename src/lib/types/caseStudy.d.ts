export interface CaseStudy {
    client: string;
    title: string;
    description: string;
    href: string;
    link: string;
    summary: string[];
    logo: Image;
    skills: string[];
    date: string;
    service: string;
    testimonial: Testimonial;
    splash: Image;
}

export interface Image {
    src: Object[],
    alt: string;
}

export interface Testimonial {
    src: Object[];
    content: string;
    author: Author;
}

export interface Author {
    name: string;
    role: string;
}