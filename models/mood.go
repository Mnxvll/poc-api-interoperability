package models

type GameMoodData struct {
	Theme        string
	ColorPalette string
	Environment  GameEnvironment
	Stats        PlayerModifiers
}

type GameEnvironment struct {
	SkyboxTexture         string
	FogDensity            float32
	WeatherParticleEffect string
}

type PlayerModifiers struct {
	StaminaDrainRate float32
	MovementSpeed    float32
}
