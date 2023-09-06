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
  icon: string;
}

export interface Image {
  src: Object[];
  alt: string;
}
