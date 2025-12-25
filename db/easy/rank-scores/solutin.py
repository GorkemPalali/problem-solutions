import pandas as pd

def order_scores(scores: pd.DataFrame) -> pd.DataFrame:
    scores_sorted = scores.sort_values('score',ascending = False)
    scores_sorted['rank']= scores_sorted['score'].rank(ascending = False, method='dense')

    return scores_sorted[['score','rank']]