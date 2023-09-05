import type { Image } from "$lib/types/sanity";

export interface ProductScreenshotData {
    left?: boolean;
    eyebrow: string;
    title: string;
    description: string;
    features: Feature[];
    image: Image;
}

export interface Feature {
    title: string;
    description: string;
    iconPath: string[];
}

export interface Image {
    src: Object[],
    alt: string
};