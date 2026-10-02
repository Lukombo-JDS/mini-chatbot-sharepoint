
class FakeEmbeddingModel:

    def embed(self,query: str) -> list[float]:

        if query == "risk":
            return [1.0, 0.0]

        elif query == "security":
            return [0.9, 0.1]

        elif query == "vacation":
            return [0.0, 1.0]

        raise ValueError("invalid entry unknown")