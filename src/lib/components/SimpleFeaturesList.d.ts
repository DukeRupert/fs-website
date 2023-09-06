export interface SimpleFeaturesListData {
    left?: boolean;
    eyebrow: string;
    title: string;
    description: string;
    features: Feature[];
}

export interface Feature {
    title: string;
    description: string;
}